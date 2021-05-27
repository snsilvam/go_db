package invoiceitem

import (
	"database/sql"
	"time"
)

//Model invoceitem
type Model struct {
	ID              uint
	invoiceheaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Slice de modelos
type Models []*Model

//Storage para crear interfaz de tx
type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, uint, Models) error
}

//Servicio de invoiceitem
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
