package invoiceitem

import "time"

//Model invoceitem
type Model struct {
	ID              uint
	invoiceheaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
type Storage interface {
	Migrate() error
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
