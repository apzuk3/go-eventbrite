package eventbrite

import (
	"fmt"

	"golang.org/x/net/context"
)

// https://www.eventbrite.com/developer/v3/endpoints/webhooks/#ebapi-id3
type WebhooksRequest struct {
	// The organization for which the webhooks will be fetched
	OrganizationID string `json:"organization_id"`
}

type WebhooksResult struct {
	Pagination Pagination `json:"pagination"`
	Webhooks   []Webhook  `json:"webhooks"`
}

// https://www.eventbrite.com/developer/v3/endpoints/webhooks/#ebapi-id5
type CreateWebhookRequest struct {
	// The target URL of the Webhook subscription
	EndpointUrl string `json:"endpoint_url"`
	// Determines what actions will trigger the webhook. If no value is sent for this param, it selects
	// order.placed, event.published, and event.unpublished by default. See below for a more complete
	// description of all available actions
	Actions string `json:"actions"`
	// The organization under which the webhook management is scoped
	OrganizationID string `json:"organization_id"`
	// The ID of the event that triggers this webhook. Leave blank for all events
	EventID string `json:"event_id"`
}

// Returns a webhook for the specified webhook as webhook
//
// https://www.eventbrite.com/developer/v3/endpoints/webhooks/#ebapi-get-webhooks-id
func (c *Client) WebhookGet(ctx context.Context, id string) (*Webhook, error) {
	res := new(Webhook)

	return res, c.getJSON(ctx, fmt.Sprintf("/webhooks/%s/", id), nil, res)
}

// Deletes the specified webhook object
//
// https://www.eventbrite.com/developer/v3/endpoints/webhooks/#ebapi-delete-webhooks-id
func (c *Client) WebhookDelete(ctx context.Context, id string) (*Webhook, error) {
	res := new(Webhook)

	return res, c.deleteJSON(ctx, fmt.Sprintf("/webhooks/%s/", id), res)
}

// Returns the list of webhook objects that belong to the authenticated user
//
// https://www.eventbrite.com/developer/v3/endpoints/webhooks/#ebapi-get-webhooks
func (c *Client) Webhooks(ctx context.Context, req *WebhooksRequest) (*WebhooksResult, error) {
	res := new(WebhooksResult)

	return res, c.getJSON(ctx, fmt.Sprintf("/webhooks/"), req, res)
}

// Creates a webhook for the authenticated user
//
// https://www.eventbrite.com/developer/v3/endpoints/webhooks/#ebapi-post-webhooks
func (c *Client) WebhookCreate(ctx context.Context, req *CreateWebhookRequest) (*Webhook, error) {
	res := new(Webhook)

	return res, c.postJSON(ctx, "/webhooks/", req, res)
}
