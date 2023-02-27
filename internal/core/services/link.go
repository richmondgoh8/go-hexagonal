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
	CreateURLData(ctx context.Context, req domain.LinkReq) error
	GetURLData(ctx context.Context, id string) (domain.Link, error)
	UpdateURLData(ctx context.Context, link domain.Link) error
}

// NewLinkSvc Takes in Interface, Return Struct
func NewLinkSvc(linkRepository ports.LinkRepository) *LinkSvc {
	return &LinkSvc{
		linkRepository: linkRepository,
	}
}

func (srv *LinkSvc) CreateURLData(ctx context.Context, req domain.LinkReq) error {
	logger.Info("Trying to insert a record to db", ctx, nil)
	return srv.linkRepository.CreateURL(ctx, req)
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
