package repo

import (
	"context"
	"database/sql"
	"errors"
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

func (db *postgres) CreateURL(ctx context.Context, req domain.LinkReq) error {
	_, err := db.postgresDB.Exec("INSERT INTO test.links (url, name) VALUES ($1, $2)", req.Url, req.Name)
	if err != nil {
		return err
	}

	return nil
}

func (db *postgres) GetURL(ctx context.Context, id string) (domain.Link, error) {
	var linkResp domain.Link
	err := db.postgresDB.Get(&linkResp, "SELECT id, url, name FROM test.links where id=$1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Link{}, errors.New("no such url was found")
		}

		return domain.Link{}, err
	}

	return linkResp, nil
}

func (db *postgres) UpdateURL(ctx context.Context, link domain.Link) error {
	sqlQuery := "UPDATE test.links SET url=$1, name=$2 where id=$3"
	_, err := db.postgresDB.ExecContext(ctx, sqlQuery, link.Url, link.Name, link.ID)

	if err != nil {
		return err
	}

	return nil
}
