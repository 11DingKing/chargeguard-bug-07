package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	result := charging.CompleteRecheck("rectified", true)
	if result.Reminder {
		w.Header().Set("X-Reminder-Queued", "true")
	}
	_ = json.NewEncoder(w).Encode(result)
}
