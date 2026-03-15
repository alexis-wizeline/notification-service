package providers

import (
	"fmt"

	"github.com/alexis-wizeline/notification-service/internals/models"
)

type InApp struct{}

func (i InApp) Deliver(n models.Notification) error {
	fmt.Printf("Sending InApp to user %d: %s\n", n.UserID, n.Message)
	return nil
}
