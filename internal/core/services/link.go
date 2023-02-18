package services

import (
	"context"
	"errors"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	"github.com/richmondgoh8/boilerplate/internal/core/ports"
	"github.com/richmondgoh8/boilerplate/pkg/logger"
)

type LinkSvc struct {
	linkRepository ports.LinkRepository
}

type LinkSvcImpl interface {
	GetURLData(ctx context.Context, id string) (domain.Link, error)
	UpdateURLData(ctx context.Context, link domain.Link) error
}

// NewLinkSvc Takes in Interface, Return Struct
func NewLinkSvc(linkRepository ports.LinkRepository) *LinkSvc {
	return &LinkSvc{
		linkRepository: linkRepository,
	}
}

func (srv *LinkSvc) GetURLData(ctx context.Context, id string) (domain.Link, error) {
	logger.Info("Retrieving Data from DB", ctx, nil)
	link, err := srv.linkRepository.GetURL(ctx, id)
	if err != nil {
		return domain.Link{}, errors.New("get link from repository has failed")
	}

	return link, nil
}

func (srv *LinkSvc) UpdateURLData(ctx context.Context, link domain.Link) error {
	_, err := srv.linkRepository.GetURL(ctx, link.ID)
	if err != nil {
		return err
	}

	return srv.linkRepository.UpdateURL(ctx, link)
}
