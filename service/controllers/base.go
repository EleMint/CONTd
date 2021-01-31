package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/EleMint/CONTd/service/contracts"
)

// InitBaseController initializes the base controller
func InitBaseController() {}

// BaseHTTPHandleFunc handles incoming requests for /
func BaseHTTPHandleFunc(w http.ResponseWriter, r *http.Request) {
	resp := &contracts.RootResponse{
		Response: contracts.NewResponse(
			http.StatusOK,
			map[string]string{},
			map[string]string{
				"UserController": "not fully implemented",
			},
		),
	}
	marsh, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal response data", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.Response.StatusCode)
	w.Write(marsh)
}
