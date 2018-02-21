package eventbrite

import "golang.org/x/net/context"

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

// Gets a refund-request for the specified refund request
//
// https://www.eventbrite.com/developer/v3/endpoints/refund_requests/#ebapi-get-refund-requests-id
func (c *Client) RefundRequest(ctx context.Context, id string) (*RefundRequest, error) {
	res := new(RefundRequest)

	return res, c.getJSON(ctx, "/refund_requests/" + id, nil, res)
}

// Update a refund-request for a specific order. Each element in items is a refund-item
//
// https://www.eventbrite.com/developer/v3/endpoints/refund_requests/#ebapi-post-refund-requests-id
func (c *Client) RefundRequestUpdate(ctx context.Context, id string, req *UpdateOrganizerRequest) (*RefundRequest, error) {
	res := new(RefundRequest)

	return res, c.postJSON(ctx, "/refund_requests/" + id, nil, res)
}

// Creates a refund-request for a specific order. Each element in items is a refund-item
//
// https://www.eventbrite.com/developer/v3/endpoints/refund_requests/#ebapi-post-refund-requests
func (c *Client) RefundRequestCreate(ctx context.Context, req *CreateRefundRequest) (*RefundRequest, error) {
	res := new(RefundRequest)

	return res, c.postJSON(ctx, "/refund_requests/", req, res)
}