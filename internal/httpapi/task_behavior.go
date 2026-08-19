package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	result := charging.CompleteRecheck("rectified", true)
	if result.State != "closed" || result.Reminder || result.Audit != "closed" {
		http.Error(w, "recheck not closed", http.StatusConflict)
		return
	}
	w.Header().Set("X-Reminder-Stopped", "true")
	_ = json.NewEncoder(w).Encode(result)
}
