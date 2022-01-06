package manager

import (
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	SqlDB() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func NewInfra(datasource string) Infra {
	conn, err := initDbResource(datasource)
	if err != nil {
		log.Fatal(err)
	}
	return &infra{db: conn}
}

func initDbResource(datasource string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", datasource)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (i *infra) SqlDB() *sqlx.DB {
	return i.db
}
