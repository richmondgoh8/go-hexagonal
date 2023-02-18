package handlers

import (
	"encoding/json"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	"github.com/richmondgoh8/boilerplate/internal/core/services"
	"net/http"
)

type TokenHandler struct {
	//linkSvc svc.LinkSvcImpl
	tokenSvc services.TokenSvcImpl
}

func NewTokenHandler(tokenSvc services.TokenSvcImpl) *TokenHandler {
	return &TokenHandler{
		tokenSvc: tokenSvc,
	}
}

func (t *TokenHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var resp domain.TokenResp
	resp.StatusCode = http.StatusOK

	token, err := t.tokenSvc.GetJWTToken(ctx)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
	}

	resp.Token = token.Token

	json.NewEncoder(w).Encode(&resp)
}
