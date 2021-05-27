package invoice

import (
	"github.com/snsilvam/go_db/pkg/invoiceheader"
	"github.com/snsilvam/go_db/pkg/invoiceitem"
)

// Se crea la estructura de la factura
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

//Se crea la interfaz para crear la factura
type Storage interface {
	Create(*Model) error
}

//Estructura del servicio de la factura
type Service struct {
	storage Storage
}

//Función que retorna un puntero de servicios
func NewService(s Storage) *Service {
	return &Service{s}
}

//Función para crear una nueva factura
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
