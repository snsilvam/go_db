package main

import (
	"log"

	"github.com/snsilvam/go_db/pkg/product"
	"github.com/snsilvam/go_db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("Erro al intentar migrar productos: %v", err)
	}
}
