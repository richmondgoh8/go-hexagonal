package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	svc "github.com/richmondgoh8/boilerplate/internal/core/services"
	"github.com/richmondgoh8/boilerplate/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type URLHandler struct {
	linkSvc svc.LinkSvcImpl
}

func NewURLHandlerImpl(linkSvc svc.LinkSvcImpl) *URLHandler {
	return &URLHandler{
		linkSvc: linkSvc,
	}
}

func (h *URLHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	linkID := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(linkID); err != nil {
		json.NewEncoder(w).Encode(&domain.SimpleResp{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	logger.Info("Start URL Handler", ctx, nil)
	resp, err := h.linkSvc.GetURLData(ctx, linkID)
	if err != nil {
		json.NewEncoder(w).Encode(&domain.SimpleResp{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	json.NewEncoder(w).Encode(&resp)
}

func (h *URLHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	r.ParseForm()
	linkID := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(linkID); err != nil {
		json.NewEncoder(w).Encode(&domain.SimpleResp{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	urlUpdateReq := &domain.Link{
		ID:   linkID,
		Url:  r.Form.Get("url"),
		Name: r.Form.Get("name"),
	}

	err := validate.Struct(urlUpdateReq)
	if err != nil {
		validatorErr := err.(validator.ValidationErrors)
		out := make([]domain.ApiError, len(validatorErr))
		for i, fe := range validatorErr {
			out[i] = domain.ApiError{Param: fe.Field(), Message: msgForTag(fe)}
		}
		json.NewEncoder(w).Encode(&domain.SimpleResp{
			Message:    out,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	err = h.linkSvc.UpdateURLData(ctx, *urlUpdateReq)
	if err != nil {
		json.NewEncoder(w).Encode(&domain.SimpleResp{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	json.NewEncoder(w).Encode(&domain.SimpleResp{
		Message:    "successfully updated records",
		StatusCode: http.StatusOK,
	})
}
