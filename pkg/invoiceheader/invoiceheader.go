package invoiceheader

//https://es.chaturbate.com/beberly_hv/?__cf_chl_captcha_tk__=28232a3e4f889d8d63f30195ec6f4c65a01da34c-1621877550-0-ARVHctfMyfzFkSH7Loksw0erq39WS0Jbq_sN0ijC5s0aJLmDLrO6vM1mzZqurRnFtwJtGbXl1tZpNibwNQ_cBhyUwlH3BrQUgHuTlWwBpNS8tOPWLQU2uiHEcTZCwbvNaeeMoorUkazD7Q8P2t7tg5gIw39-ituqCGxs4AEd8_zBb62l0FZ5Qic-7e6XE5ygXzE6bbjH6C_XQqOeoMnHqqPCqqsIVmAjAcquV5RKh3t72d494AiqkG_-OaxwTgLapafxs57QjgHcUmmmjTi1b3ZZlbQCc-RJ9fpk9eitM5cPoDG7vGnTlhCMhOUEqROlQUlZ50MYTH2hqgNEGk475qzZPj2_6MwCklsb3CpWYz_MzQ3XT7ySsDroHWaUplSJ_maPswMxDrswHPvD6XrlopBiThKzm4p3USE8s6I_DoczU-zMFYc6HiumNjVE4kM7AE6R77IRXbrHi-UMpX4OBDCE_85BqPVI9pKZ9fh2RAcnFe507XeqD4eFg7GziRSgPRrgDZuP51XV82aCsRNSg32ii3_QPrLwkNBWv-vFzjMeuGUKohyvvoQ9skQ19dJy7rl9Nq5yiiBPOWa8Beet5hQBywLeRiUnvEzngsrZbiJZ3_iWWEAev31CEaipe84vxCcE2fHolVp9PXaZqVk1uAfCQqBvfowSOAgtIyMx-UJr
import (
	"database/sql"
	"time"
)

//Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Slice de modelos
type Models []*Model

//Interfaz para implementar la tx
type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
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
