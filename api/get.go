package api

import (
	"net/http"
	"encoding/json"

	"github.com/vlamug/scfg/storage"
	"github.com/vlamug/scfg/request"
)

func GetHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		req := request.GetRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cfg := storage.Get(req)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cfg.CSet))
	}
}
