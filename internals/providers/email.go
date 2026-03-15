package providers

import (
	"fmt"

	"github.com/alexis-wizeline/notification-service/internals/models"
)

type EmailProvider struct{}

func (e EmailProvider) Deliver(n models.Notification) error {
	fmt.Printf("Sending Email to user %d: %s\n", n.UserID, n.Message)
	return nil
}
