package version

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// RegisterHandler registers version handler to /version
func RegisterHandler(mu *http.ServeMux, logger *zap.Logger) {
	info := Get()
	json, err := json.Marshal(info)
	if err != nil {
		logger.Fatal("Could not get Jaeger version", zap.Error(err))
	}
	mu.HandleFunc("/version", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
		w.Write(json)
	})
}
