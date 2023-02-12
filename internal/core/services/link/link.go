package svc

import (
	"context"
	"errors"

	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	"github.com/richmondgoh8/boilerplate/internal/core/ports"
)

type LinkSvc struct {
	linkRepository ports.LinkRepository
}

// Takes in Interface, Return Struct
func NewLinkSvc(linkRepository ports.LinkRepository) *LinkSvc {
	return &LinkSvc{
		linkRepository: linkRepository,
	}
}

func (srv *LinkSvc) Get(ctx context.Context, id string) (domain.Link, error) {
	link, err := srv.linkRepository.Get(ctx, id)
	if err != nil {
		return domain.Link{}, errors.New("get link from repository has failed")
	}

	return link, nil
}
