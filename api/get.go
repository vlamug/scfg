package api

import (
	"net/http"
	"encoding/json"

	"github.com/vlamug/scfg/request"
	"github.com/vlamug/scfg/loader"
)

// GetHandler loads config by key
func GetHandler(loaderService *loader.Loader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		req := request.GetRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cfg := loaderService.Load(req)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cfg))
	}
}
