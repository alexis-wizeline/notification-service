package api

import (
	"encoding/json"
	"net/http"

	"github.com/alexis-wizeline/notification-service/internals/models"
)

type Handler struct {
	Queue chan<- models.Notification
}

func (h Handler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var n models.Notification

	err := json.NewDecoder(r.Body).Decode(&n)

	if err != nil {
		http.Error(w, "invalid request", 400)
	}

	n.Status = models.Pendign

	h.Queue <- n

	w.WriteHeader(http.StatusAccepted)
}
