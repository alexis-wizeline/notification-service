package main

import (
	"net/http"

	"github.com/alexis-wizeline/notification-service/internals/api"
	"github.com/alexis-wizeline/notification-service/internals/models"
	"github.com/alexis-wizeline/notification-service/internals/providers"
	"github.com/alexis-wizeline/notification-service/internals/service"
	"github.com/alexis-wizeline/notification-service/internals/worker"
)

func main() {

	dispatcher := service.Dispatcher{
		Email: providers.EmailProvider{},
		SMS:   providers.SMSProvider{},
		InApp: providers.InApp{},
	}

	queue := make(chan models.Notification, 100)
	w := worker.Worker{
		Queue:      queue,
		Dispatcher: dispatcher,
	}

	go w.Start()

	handler := api.Handler{
		Queue: queue,
	}

	http.HandleFunc("/notifications", handler.CreateNotification)

	http.ListenAndServe(":8080", nil)
}
