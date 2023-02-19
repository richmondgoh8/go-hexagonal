package handlers

import (
	"encoding/json"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	"net/http"
)

func ReturnAPIErr(w http.ResponseWriter, errMsg interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&domain.SimpleResp{
		Message:    errMsg,
		StatusCode: statusCode,
	})
}
