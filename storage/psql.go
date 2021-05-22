package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		obvervaciones VARCHAR(100),
		price INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
		)`
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
	stmt, err := p.db.Prepare(psqlCreateProduct)
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
