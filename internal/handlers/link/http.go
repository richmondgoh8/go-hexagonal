package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/richmondgoh8/boilerplate/internal/core/ports"
)

type URLHandler struct {
	linkRepository ports.LinkRepository
}

type URLHandlerImpl interface {
	Get(http.ResponseWriter, *http.Request)
}

func NewURLHandlerImpl(linkRepository ports.LinkRepository) *URLHandler {
	return &URLHandler{
		linkRepository: linkRepository,
	}
}

func (h *URLHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	linkID := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(linkID); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	resp, err := h.linkRepository.Get(ctx, linkID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}
