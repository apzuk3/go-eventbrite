package eventbrite

import "golang.org/x/net/context"

// NotificationsResult is the response structure fornotifications
type NotificationsResult struct {
	Notifications []Notification
	Pagination    Pagination
}

// Notification is the representation of something that Eventbrite has notified to its users.
//
// see @https://www.eventbrite.com/developer/v3/response_formats/notification/#ebapi-std:format-notification
type Notification struct {
	// Notification ID
	ID string `json:"notification_id"`
	// The title of the notification
	Title string `json:"title"`
	// It is the secondary text of the notification
	Body string `json:"body"`
	// It is the call to action title associated to the notification
	Cta string `json:"cta"`
	// It is a link or a deep link associated to the notification
	Url string `json:"url"`
	// It is the date when the notification was created
	Created Date `json:"created"`
	// It is the channel (group) of the notification. The available channel ids
	// are: ticket_reminders, my_events, social, discovery and eventbrite
	ChannelId string `json:"channel_id"`
	// The image logo for the event (optional)
	ImageId Image `json:"image_id"`
}

// Notifications gets a paginated response of notification objects for a determined user.
//
// https://www.eventbrite.com/developer/v3/endpoints/notifications/#ebapi-get-users-me-notifications
func (c *Client) Notifications(ctx context.Context) (*NotificationsResult, error) {
	res := new(NotificationsResult)

	return res, c.getJSON(ctx, "/notificaitons", nil, res)
}
