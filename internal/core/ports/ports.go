package ports

import (
	"context"

	"github.com/richmondgoh8/boilerplate/internal/core/domain"
)

// LinkRepository contain interface definitions to communicate with Actors (Things Existing out of Core)
// Connects with repositories folder
type LinkRepository interface {
	GetURL(ctx context.Context, id string) (domain.Link, error)
	UpdateURL(ctx context.Context, link domain.Link) error
}
