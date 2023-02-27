package handlers

import (
	"encoding/json"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	"github.com/richmondgoh8/boilerplate/internal/core/services"
	"net/http"
)

type TokenHandler struct {
	tokenSvc services.TokenSvcImpl
}

func NewTokenHandler(tokenSvc services.TokenSvcImpl) *TokenHandler {
	return &TokenHandler{
		tokenSvc: tokenSvc,
	}
}

// swagger:route GET /token Token token_id
//
// # Get JWT Token
//
// This will generate a JWT Token to use other endpoints
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: body:TokenResp
func (t *TokenHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var resp domain.TokenResp
	resp.StatusCode = http.StatusOK

	token, err := t.tokenSvc.GetJWTToken(ctx)
	if err != nil {
		ReturnAPIErr(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Token = token.Token

	json.NewEncoder(w).Encode(&resp)
}
