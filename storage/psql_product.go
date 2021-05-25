package storage

import (
	"database/sql"
	"fmt"

	"github.com/snsilvam/go_db/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	//Constante para crear tablas
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		obvervaciones VARCHAR(100),
		price INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
		)`
	//Constante para insertar productos
	psqlCreateProduct = `INSERT INTO products(name, obvervaciones, price, create_at) VALUES($1, $2, $3, $4) RETURNING id`
	//Constante para consultar productos
	psqlGetAllProduct = `SELECT id, name, obvervaciones, price, create_at, update_at FROM products`
	//Consulta de fila por id
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	//Actualizar producto
	psqlUpdateProduct = `UPDATE products SET name = $1, obvervaciones = $2, 
	price = $3, update_at = $4 WHERE id = $5`
	//Eliminar un producto
	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

//PsqlProduct se usa para trabajar con postgre en producto
type PsqlProduct struct {
	db *sql.DB
}

//Retorna un nuevo pointerofpsqlproduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

//Migrate implementa la interfaz product.storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migración realizada")
	return nil
}

//Implementa la interfaz de product storage para crear un nuevo producto
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("Se creo el producto")
	return nil
}

//Implementación de la estructura que nos permitira ver los datos de la db
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))

}

//Consultar producto por fila
func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

//Actualizar producto
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}
	//Controlar cuando un id no existe
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("No se encontro id ingresado: %d", m.ID)
	}
	fmt.Println("Actualizanción exitosa !")
	return nil
}

//Eliminar un producto
func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Producto eliminado")
	return nil

}
