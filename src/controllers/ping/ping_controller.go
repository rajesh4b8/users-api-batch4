package ping

import (
	"encoding/json"
	"net/http"

	"github.com/rajesh4b8/users-api-batch4/src/logger"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Info("Ping request")

	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Pong!")
}
