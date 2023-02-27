package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	svc "github.com/richmondgoh8/boilerplate/internal/core/services"
	"github.com/richmondgoh8/boilerplate/pkg/logger"
	custommiddleware "github.com/richmondgoh8/boilerplate/pkg/middleware"
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

// swagger:route GET /url/{id} URL url_get_id
//
// # Get URL Mapping
//
// This will get the key value pair of URL
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//	    Parameters:
//	      + name: id
//	        in: path
//	        description: maximum numnber of results to return
//	        required: false
//	        type: integer
//	        format: int32
//
//		Responses:
//		  200: body:Link
func (h *URLHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, ok := ctx.Value("jwt_claims").(*custommiddleware.Claims)
	if !ok {
		ReturnAPIErr(w, "claims not found in context", http.StatusInternalServerError)
		return
	}
	if claims.Role != "admin" {
		ReturnAPIErr(w, "must be an admin to continue", http.StatusInternalServerError)
		return
	}

	linkID := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(linkID); err != nil {
		ReturnAPIErr(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("Start URL Handler", ctx, nil)
	resp, err := h.linkSvc.GetURLData(ctx, linkID)
	if err != nil {
		ReturnAPIErr(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&resp)
}

// swagger:route PUT /url/{id} URL url_update_id
//
// # Update URL Mapping using Form Data
//
// This will update the key value pair of URL ID in DB
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//	    Parameters:
//	      + name: id
//	        in: path
//	        type: integer
//	      + name: url
//	        in: query
//	        type: string
//	      + name: name
//	        in: query
//	        type: string
//
//		Responses:
//		  200: body:SimpleResp
func (h *URLHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	// 2mb
	r.ParseMultipartForm(1 << 20)
	linkID := chi.URLParam(r, "id")
	if _, err := strconv.Atoi(linkID); err != nil {
		ReturnAPIErr(w, err.Error(), http.StatusBadRequest)
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
		ReturnAPIErr(w, out, http.StatusBadRequest)
		return
	}

	err = h.linkSvc.UpdateURLData(ctx, *urlUpdateReq)
	if err != nil {
		ReturnAPIErr(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&domain.SimpleResp{
		Message:    "successfully updated records",
		StatusCode: http.StatusOK,
	})
}

// swagger:route POST /url URL url_create_id
//
// # Create URL Mapping
//
// This will generate the key value pair of URL
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//	    Parameters:
//	      + name: input
//	        in: body
//	        type: LinkReq
//
//		Responses:
//		  200: body:SimpleResp
func (h *URLHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	var linkReq domain.LinkReq
	if err := json.NewDecoder(r.Body).Decode(&linkReq); err != nil {
		ReturnAPIErr(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := validate.Struct(linkReq)
	if err != nil {
		validatorErr := err.(validator.ValidationErrors)
		out := make([]domain.ApiError, len(validatorErr))
		for i, fe := range validatorErr {
			out[i] = domain.ApiError{Param: fe.Field(), Message: msgForTag(fe)}
		}
		ReturnAPIErr(w, out, http.StatusBadRequest)
		return
	}

	err = h.linkSvc.CreateURLData(ctx, linkReq)
	if err != nil {
		ReturnAPIErr(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&domain.SimpleResp{
		Message:    "successfully created records",
		StatusCode: http.StatusOK,
	})
}
