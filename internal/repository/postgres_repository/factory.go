package postgres_repository

import (
	"database/sql"
	"fio_finder/internal/repository"
	"github.com/jmoiron/sqlx"
)

func CreatePersonPostgresRepository(db *sql.DB) repository.PersonRepository {
	dbx := sqlx.NewDb(db, "pgx")

	return NewPersonPostgresRepository(dbx)
}
