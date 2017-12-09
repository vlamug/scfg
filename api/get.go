package api

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/vlamug/scfg/request"
	"github.com/vlamug/scfg/loader"
	"github.com/vlamug/scfg/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

// GetHandler loads config by key
func GetHandler(loaderService *loader.Loader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time_start := time.Now()

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

		metrics.ResponseTime.With(prometheus.Labels{"handler": "get"}).Set(float64(time.Since(time_start))/float64(time.Millisecond))
	}
}
