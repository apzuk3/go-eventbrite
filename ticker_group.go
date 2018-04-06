package eventbrite

import "golang.org/x/net/context"

// The Ticket Group object allows the users to group an arbitrary number of ticket_class
//
// https://www.eventbrite.com/developer/v3/response_formats/ticket_group/#ebapi-std:format-ticket_group
type TicketGroup struct {
	ID string `json:"id"`
	// Name of the ticket group. If it is greater than 20 characters will be truncated automatically
	Name string `json:"name"`
	// The status of the ticket group. One of transfer, live, deleted or archived. By default is live is not specified
	Status string `json:"status"`
	// A dictionary with the event ids and ticket ids, in the form {event_id: [ticket_id_1, ticket_id_2]}
	//
	// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-std:format-dict
	EventTicketIds interface{} `json:"event_ticket_ids"`
	// The list of ticket class that includes the id, event_id, sales_channels, variants and name.
	// Simplified version of ticket_class. By default is empty, unless the tickets expansion is used
	Tickets []TicketClass `json:"tickets"`
}

// CreateTicketGroupRequest is the request structure to create a new ticket group
//
// https://www.eventbrite.com/developer/v3/endpoints/ticket_groups/#ebapi-id3
type CreateTicketGroupRequest struct {
	// Name of ticket group
	Name string `json:"ticket_group.name" validate:"required"`
	// The status of ticket group. Valid choices are: live, deleted, or archived
	Status string `json:"ticket_group.status"`
	// (‘IDs of tickets by event id for this ticket group. In the format “{“event_id”: [“ticket_class_id”, “ticket_class_id”]}”.’,)
	//
	// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-dictionary
	Ids map[string]interface{} `json:"ticket_group.event_ticket_ids"`
}

// UpdateTicketGroupRequest is the request structure to update ticket group
//
// https://www.eventbrite.com/developer/v3/endpoints/ticket_groups/#ebapi-id5
type UpdateTicketGroupRequest struct {
	// Name of ticket group
	Name string `json:"ticket_group.name"`
	// The status of ticket group. Valid choices are: live, deleted, or archived
	Status string `json:"ticket_group.status"`
	// (‘IDs of tickets by event id for this ticket group. In the format “{“event_id”: [“ticket_class_id”, “ticket_class_id”]}”.’,)
	//
	// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-dictionary
	Ids map[string]interface{} `json:"ticket_group.event_ticket_ids"`
}

// TicketGroupGet returns the ticket_group with the specified :ticket_group_id
//
// https://www.eventbrite.com/developer/v3/endpoints/ticket_groups/#ebapi-get-ticket-groups-ticket-group-id
func (c *Client) TicketGroupGet(ctx context.Context, id string) (*TicketGroup, error) {
	res := new(TicketGroup)

	return res, c.getJSON(ctx, "/ticket_groups/"+id, nil, res)
}

// TicketGroupGet deletes the ticket_group with the specified :ticket_group_id. The status of the ticket group is changed to deleted.
//
// https://www.eventbrite.com/developer/v3/endpoints/ticket_groups/#ebapi-delete-ticket-groups-ticket-group-id
func (c *Client) TicketGroupDelete(ctx context.Context, id string) (interface{}, error) {
	var res interface{}
	return res, c.deleteJSON(ctx, "/ticket_groups/"+id, &res)
}

// TicketGroupGet creates a ticket group and returns the created ticket_group. Only up to 200 live ticket groups may be created;
// those with archived or deleted status are not taken into account
//
// https://www.eventbrite.com/developer/v3/endpoints/ticket_groups/#ebapi-post-ticket-groups
func (c *Client) TicketGroupCreate(ctx context.Context, id string, req *CreateTicketGroupRequest) (*TicketGroup, error) {
	res := new(TicketGroup)

	return res, c.postJSON(ctx, "/ticket_groups/"+id, req, &res)
}

// TicketGroupGet updates the ticket group with the specified :ticket_group_id. Returns the updated ticket_group
//
// https://www.eventbrite.com/developer/v3/endpoints/ticket_groups/#ebapi-post-ticket-groups-ticket-group-id
func (c *Client) TicketGroupUpdate(ctx context.Context, id string, req *UpdateTicketGroupRequest) (*TicketGroup, error) {
	res := new(TicketGroup)

	return res, c.postJSON(ctx, "/ticket_groups/"+id, req, &res)
}
