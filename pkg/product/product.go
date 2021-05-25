package product

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrIDNotFound = errors.New("No se encontro producto ingresado")
)

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//Funcion para ver lo datos de la db de forma ordenada
func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s |  %10s",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

//Models slice of model
type Models []*Model
type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}

//Servicio de producto
type Service struct {
	storage Storage
}

//Nuevo servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

//Utilizamos esta función para migrar el producto
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

//Crear un nuevo producto
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

//Consultar todos los productos
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

//Consulta por id
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

//Actualizar un producto de la db
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	m.UpdatedAt = time.Now()

	return s.storage.Update(m)
}

//Eliminar un producto
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
