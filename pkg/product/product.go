package product

import "time"

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//Models slice of model
type Models []*Model
type Storage interface {
	Migrate() error

	/* Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error */
}

//Servicio de producto
type Service struct {
	storage Storage
}

//Nuevo servicio
func NewService(s Storage) *Service {
	return &Service{s}
}
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
