package providers

import "github.com/alexis-wizeline/notification-service/internals/models"

type Provider interface {
	Deliver(n models.Notification) error
}
