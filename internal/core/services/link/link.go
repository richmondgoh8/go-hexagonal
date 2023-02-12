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

// NewLinkSvc Takes in Interface, Return Struct
func NewLinkSvc(linkRepository ports.LinkRepository) *LinkSvc {
	return &LinkSvc{
		linkRepository: linkRepository,
	}
}

func (srv *LinkSvc) GetURL(ctx context.Context, id string) (domain.Link, error) {
	link, err := srv.linkRepository.GetURL(ctx, id)
	if err != nil {
		return domain.Link{}, errors.New("get link from repository has failed")
	}

	return link, nil
}

func (srv *LinkSvc) UpdateURL(ctx context.Context, link domain.Link) error {
	_, err := srv.linkRepository.GetURL(ctx, link.ID)
	if err != nil {
		return err
	}

	return srv.linkRepository.UpdateURL(ctx, link)
}
