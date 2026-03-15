package service

import (
	"errors"

	"github.com/alexis-wizeline/notification-service/internals/models"
	"github.com/alexis-wizeline/notification-service/internals/providers"
)

type Dispatcher struct {
	Email providers.Provider
	SMS   providers.Provider
	InApp providers.Provider
}

func (d Dispatcher) Dispatch(n models.Notification) error {
	switch n.Channel {
	case "email":
		return d.Email.Deliver(n)
	case "sms":
		return d.SMS.Deliver(n)
	case "inapp":
		return d.InApp.Deliver(n)
	default:
		return errors.New("unknow provider")
	}
}
