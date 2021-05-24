package invoiceheader

import "time"

//Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Storage interface {
	Migrate() error
}

//Servicio de invoice header
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
