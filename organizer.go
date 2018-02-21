package eventbrite

import (
	"fmt"

	"golang.org/x/net/context"
)

type CreateOrganizerRequest struct {
	Name string `json:"organizer.name" validate:"required"`
	Description string `json:"organizer.description.html"`
	LongDescription string `json:"organizer.long_description.html"`
	LogoId string `json:"organizer.logo.id"`
	Website string `json:"organizer.website"`
	Twitter string `json:"organizer.twitter"`
	Facebook string `json:"organizer.facebook"`
	Instagram string `json:"organizer.instagram"`
}

type UpdateOrganizerRequest struct {
	Name string `json:"organizer.name"`
	Description string `json:"organizer.description.html"`
	LongDescription string `json:"organizer.long_description.html"`
	LogoId string `json:"organizer.logo.id"`
	Website string `json:"organizer.website"`
	Twitter string `json:"organizer.twitter"`
	Facebook string `json:"organizer.facebook"`
	Instagram string `json:"organizer.instagram"`
}

type OrganizerEventsRequest struct {

	// Only return events with a specific status set. This should be a comma delimited string of status.
	// Valid status: all, draft, live, canceled, started, ended.
	Status string `json:"status"`

	// How to order the results (Valid choices are: start_asc, start_desc, created_asc, or created_desc)
	OrderBy string `json:"order_by"`

	// Only return events with start dates after the given date
	StartDateRangeStart string `json:"start_date.range_start"`

	// Only return events with start dates after the given date
	StartDateRangeEnd string `json:"start_date.range_end"`

	// Only show public events even if viewing your own events.
	PublicOnly bool `json:"only_public"`
}

type OrganizerEventsResult struct {
	Events []Event `json:"events"`
	Pagination Pagination `json:"pagination"`
}

// Makes a new organizer. Returns the organizer
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-post-organizers
func (c *Client) OrganizerCreate(ctx context.Context, req *CreateOrganizerRequest) (*Organizer, error) {
	resp := new(Organizer)

	return resp, c.postJSON(ctx, "/organizers/", req, resp)
}

// Gets an organizer by ID as organizer.
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-get-organizers-id
func (c *Client) OrganizerGet(ctx context.Context, id string) (*Organizer, error) {
	resp := new(Organizer)

	return resp, c.postJSON(ctx, "/organizers/" + id, nil, resp)
}

// Updates an organizer and returns it as as organizer.
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-post-organizers
func (c *Client) OrganizerUpdate(ctx context.Context, id string, req *UpdateOrganizerRequest) (*Organizer, error) {
	resp := new(Organizer)

	return resp, c.postJSON(ctx, "/organizers/" + id, req, resp)
}

// Gets events of the organizer.
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-get-organizers-id-events
func (c *Client) OrganizerGetEvents(ctx context.Context, id string, req *OrganizerEventsRequest) (*OrganizerEventsResult, error) {
	resp := new(OrganizerEventsResult)

	return resp, c.getJSON(ctx, fmt.Sprintf("/organizers/%s/events/", id), req, resp)
}