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
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}
