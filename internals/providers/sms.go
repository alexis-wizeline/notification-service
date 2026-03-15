package providers

import (
	"fmt"

	"github.com/alexis-wizeline/notification-service/internals/models"
)

type SMSProvider struct{}

func (s SMSProvider) Deliver(n models.Notification) error {
	fmt.Printf("Sending SMS to user %d: %s\n", n.UserID, n.Message)
	return nil
}
