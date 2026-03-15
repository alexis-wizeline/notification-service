package worker

import (
	"log"

	"github.com/alexis-wizeline/notification-service/internals/models"
	"github.com/alexis-wizeline/notification-service/internals/service"
)

type Worker struct {
	Queue      chan models.Notification
	Dispatcher service.Dispatcher
}

func (w Worker) Start() {
	for notification := range w.Queue {

		err := w.Dispatcher.Dispatch(notification)

		if err != nil {
			log.Println("Failed:", err)
			continue
		}

		log.Println("Notification sent: ", notification.ID)
	}
}
