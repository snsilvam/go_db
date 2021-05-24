package main

import (
	"fmt"
	"log"

	"github.com/snsilvam/go_db/pkg/product"
	"github.com/snsilvam/go_db/storage"
)

func main() {
	storage.NewPostgresDB()
	//*******ESTO CREA LA TABLA PRODUCT
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// if err := serviceProduct.Migrate(); err != nil {
	// 	log.Fatalf("Erro al intentar migrar productos: %v", err)
	// }
	//*******ESTO CREA LA TABLA HEADER
	// storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	// if err := serviceInvoiceHeader.Migrate(); err != nil {
	// 	log.Fatalf("invoiceHeader.Migrate: %v", err)
	// }
	//*******ESTO CREA LA TABLA ITEM
	// storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	// serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	// if err := serviceInvoiceItem.Migrate(); err != nil {
	// 	log.Fatalf("invoiceItem.Migrate: %v", err)
	// }
	//******Agregar un producto
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	//Modelo del producto que ingresaremos en la tabla
	// m := &product.Model{
	// 	Name:         "Druid",
	// 	Price:        1000,
	// 	Observations: "Es el guardian de la naturaleza",
	// }
	//*********Consultar una tabla
	// ms, err := serviceProduct.GetAll()
	// if err != nil {
	// 	log.Fatalf("Error.verTabla: %v", err)
	// }
	// fmt.Println(ms)
	//**Consultar un producto por id
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m, err := serviceProduct.GetByID(3)
	if err != nil {
		log.Fatalf("product.GetByID: %v", err)
	}
	fmt.Println(m)
}
