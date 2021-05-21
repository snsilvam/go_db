package storage

import (
	// ...
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // paquete del driver para PostgreSQL
)

var (
	db   *sql.DB   // estructura db gestiona un pool de conexiones activas e inactivas
	once sync.Once // estructura Once que permite ejecutar una única vez (Singleton)
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		// El primer argumento de Open es el nombre del driver ("postgres")
		// y el segundo argumento es la cadena de conexión, donde se coloca
		// las credenciales de acceso a la BD
		db, err = sql.Open("postgres", "postgres://edteam:edteam@localhost:7530/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// Pool retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}
