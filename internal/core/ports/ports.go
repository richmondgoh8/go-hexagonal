package ports

import (
	"context"

	domain "github.com/richmondgoh8/boilerplate/internal/core/domain"
)

// contain interface definitions to communicate with Actors (Things Existing out of Core)
// Actors can refer to DB, Redis or Other Storage Mediums
type LinkRepository interface {
	Get(ctx context.Context, id string) (domain.Link, error)
}
