package components

import (
	"database/sql"

	"github.com/AbiXnash/theta-api/internals/service"
)

type Components struct {
	DB     *sql.DB
	Status *service.StatusService
}

func New() *Components {
	return &Components{}
}
