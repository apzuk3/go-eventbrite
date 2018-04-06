package eventbrite

import (
	"fmt"
	"golang.org/x/net/context"
)

// Order is an object representing an order made against Eventbrite for one or more ticket classes
// Order objects are considered private and are only available to the event owner and the person who placed the order
//
// https://www.eventbrite.com/developer/v3/response_formats/order/#ebapi-std:format-order
type Order struct {
	// When the attendee was created (order placed)
	Created DateTime `json:"created"`
	// When the attendee was last changed
	Changed DateTime `json:"changed"`
	// The ticket buyer’s name. Use this in preference to
	// first_name/last_name if possible for forward compatibility with non-Western names
	Name string `json:"name"`
	// The ticket buyer’s first name
	FirstName string `json:"first_name"`
	// The ticket buyer’s last name
	LastName string `json:"last_name"`
	// The ticket buyer’s email address
	Email string `json:"email"`
	// Cost breakdown for this order
	Costs OrderCosts `json:"costs"`
	// The event this order is against
	Event Event `json:"event"`
	// Refund request on this order
	RefundRequests RefundRequest `json:"refund_requests"`
	//Attendees on this order
	Attendees []Attendee `json:"attendees"`
	// The event id this order is against
	EventID string `json:"event_id"`
	// The time remaining to complete this order (in seconds)
	TimeRemaining int `json:"time_remaining"`
}

// OrderCosts contains a breakdown of the order’s costs
//
// https://www.eventbrite.com/developer/v3/response_formats/order/#ebapi-order-costs
type OrderCosts struct {
	// The total amount the buyer was charged
	Gross Currency `json:"gross"`
	// The portion of gross taken by Eventbrite as a management fee
	EventbriteFee Currency `json:"eventbrite_fee"`
	// The portion of gross taken by the payment processor
	PaymentFee Currency `json:"payment_fee"`
	// The portion of gross allocated for tax (but passed onto the organizer)
	Tex Currency `json:"tax"`
}

// OrderGet gets an order by ID an order object
//
// https://www.eventbrite.com/developer/v3/endpoints/orders/#ebapi-orders
func (c *Client) OrderGet(ctx context.Context, id string) (*Order, error) {
	o := new(Order)

	return o, c.getJSON(ctx, fmt.Sprintf("/orders/%s/", id), nil, o)
}
