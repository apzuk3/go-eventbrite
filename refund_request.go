package eventbrite

import "golang.org/x/net/context"

// RefundRequest contains a refund request of the order
//
// https://www.eventbrite.com/developer/v3/response_formats/order/#ebapi-std:format-refund-request
type RefundRequest struct {
	// The email used to create the refund request
	FromEmail string `json:"from_email"`
	// The name used to create the refund request
	FromName string `json:"from_name"`
	// The actual status of the refund request
	Status string `json:"status"`
	// The message associated with the refund request
	Message string `json:"message"`
	// The code of the refund request’s reason
	Reason string `json:"reason"`
	// The last message associated with the last status of the refund request
	LastMessage string `json:"last_message"`
	// The last code of the refund request’s reason
	LastReason string `json:"last_reason"`
	// The items of the refund request
	Items []RefundItem `json:"items"`
}

// RefundItem contains a refund item
//
// https://www.eventbrite.com/developer/v3/response_formats/order/#ebapi-refund-item
type RefundItem struct {
	// The event of this item
	EventID string `json:"event_id"`
	// The order of this item. Consider that this field can be null
	OrderID string `json:"order_id"`
	// the item type order for full refund, attendee for partial refund an
	// attendee or merchandise for partial refund a merchandise
	ItemType string `json:"item_type"`
	// The quantity requested for this item. If the item_type is order, quantity_requested is always 1.
	// if the item_type is attendee or merchandise, then the quantity_requested shows how
	// many items were requested
	QuantityRequested int `json:"quantity_requested"`
	// The total amount requested for this item.
	AmountRequested Currency `json:"amount_requested"`
}

// CreateRefundRequest is the request structure to create a refund request
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/refund_requests/#ebapi-id1
type CreateRefundRequest struct {
	// The email used to create the refund request
	FromEmail string `json:"from_email" validate:"required"`
	// The name used to create the refund request
	FromName string `json:"from_name" validate:"required"`
	// The items of the refund request
	Items []RefundItem `json:"items" validate:"required"`
	// The message associated with the refund request
	Message string `json:"message" validate:"required"`
	// The code of the refund request’s reason
	Reason string `json:"reason" validate:"required"`
}

// CreateRefundRequest is the request structure to update refund request
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/refund_requests/#ebapi-id3
type UpdateRefundRequest struct {
	// The email used to create the refund request
	FromEmail string `json:"from_email" validate:"required"`
	// The name used to create the refund request
	FromName string `json:"from_name" validate:"required"`
	// The items of the refund request
	Items []RefundItem `json:"items" validate:"required"`
	// The message associated with the refund request
	Message string `json:"message" validate:"required"`
	// The code of the refund request’s reason
	Reason string `json:"reason" validate:"required"`
}

// RefundRequest gets a refund-request for the specified refund request
//
// https://www.eventbrite.com/developer/v3/endpoints/refund_requests/#ebapi-get-refund-requests-id
func (c *Client) RefundRequest(ctx context.Context, id string) (*RefundRequest, error) {
	res := new(RefundRequest)

	return res, c.getJSON(ctx, "/refund_requests/"+id, nil, res)
}

// RefundRequestUpdate updates a refund-request for a specific order. Each element in items is a refund-item
//
// https://www.eventbrite.com/developer/v3/endpoints/refund_requests/#ebapi-post-refund-requests-id
func (c *Client) RefundRequestUpdate(ctx context.Context, id string, req *UpdateOrganizerRequest) (*RefundRequest, error) {
	res := new(RefundRequest)

	return res, c.postJSON(ctx, "/refund_requests/"+id, nil, res)
}

// RefundRequestCreate creates a refund-request for a specific order. Each element in items is a refund-item
//
// https://www.eventbrite.com/developer/v3/endpoints/refund_requests/#ebapi-post-refund-requests
func (c *Client) RefundRequestCreate(ctx context.Context, req *CreateRefundRequest) (*RefundRequest, error) {
	res := new(RefundRequest)

	return res, c.postJSON(ctx, "/refund_requests/", req, res)
}
