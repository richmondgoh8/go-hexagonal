package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
)

type postgres struct {
	postgresDB *sqlx.DB
}

func NewPostgresInstance(postgresDB *sqlx.DB) *postgres {
	return &postgres{
		postgresDB: postgresDB,
	}
}

func (db *postgres) Get(ctx context.Context, id string) (domain.Link, error) {
	var linkResp domain.Link
	err := db.postgresDB.Get(&linkResp, "SELECT id, url, name FROM test.links where id=$1", id)
	if err != nil {
		return domain.Link{}, err
	}

	return linkResp, nil
}
