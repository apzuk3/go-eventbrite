package eventbrite

import (
	"fmt"

	"golang.org/x/net/context"
)

// CreateOrganizerRequest is the request structure for creating a new organizer
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/organizers/#ebapi-parameters
type CreateOrganizerRequest struct {
	// The name of the organizer
	Name string `json:"organizer.name" validate:"required"`
	// The description of the organizer
	Description string `json:"organizer.description.html"`
	// The long description of the organizer
	LongDescription string `json:"organizer.long_description.html"`
	// The logo id of the organizer
	LogoID string `json:"organizer.logo.id"`
	// The website for the organizer
	Website string `json:"organizer.website"`
	// The Twitter handle for the organizer
	Twitter string `json:"organizer.twitter"`
	// The Facebook URL ID for the organizer
	Facebook string `json:"organizer.facebook"`
	// The Instagram numeric ID for the organizer
	Instagram string `json:"organizer.instagram"`
}

// UpdateOrganizerRequest is the request structure for updating an organizer
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/organizers/#ebapi-id3
type UpdateOrganizerRequest struct {
	// The name of the organizer
	Name string `json:"organizer.name"`
	// The description of the organizer
	Description string `json:"organizer.description.html"`
	// The long description of the organizer
	LongDescription string `json:"organizer.long_description.html"`
	// The logo id of the organizer
	LogoId string `json:"organizer.logo.id"`
	// The website for the organizer
	Website string `json:"organizer.website"`
	// The Twitter handle for the organizer
	Twitter string `json:"organizer.twitter"`
	// The Facebook URL ID for the organizer
	Facebook string `json:"organizer.facebook"`
	// The Instagram numeric ID for the organizer
	Instagram string `json:"organizer.instagram"`
}

// OrganizerEventsRequest is the request structure to get organizer events
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/organizers/#ebapi-id6
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

// OrganizerEventsResult is the response structure for organizer events request
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/organizers/#ebapi-get-organizers-id-events
type OrganizerEventsResult struct {
	Events     []Event    `json:"events"`
	Pagination Pagination `json:"pagination"`
}

// OrganizerCreate makes a new organizer. Returns the organizer
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-post-organizers
func (c *Client) OrganizerCreate(ctx context.Context, req *CreateOrganizerRequest) (*Organizer, error) {
	resp := new(Organizer)

	return resp, c.postJSON(ctx, "/organizers/", req, resp)
}

// OrganizerCreate gets an organizer by ID as organizer.
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-get-organizers-id
func (c *Client) OrganizerGet(ctx context.Context, id string) (*Organizer, error) {
	resp := new(Organizer)

	return resp, c.postJSON(ctx, "/organizers/"+id, nil, resp)
}

// OrganizerCreate updates an organizer and returns it as as organizer.
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-post-organizers
func (c *Client) OrganizerUpdate(ctx context.Context, id string, req *UpdateOrganizerRequest) (*Organizer, error) {
	resp := new(Organizer)

	return resp, c.postJSON(ctx, "/organizers/"+id, req, resp)
}

// OrganizerCreate gets events of the organizer.
//
// https://www.eventbrite.com/developer/v3/endpoints/organizers/#ebapi-get-organizers-id-events
func (c *Client) OrganizerGetEvents(ctx context.Context, id string, req *OrganizerEventsRequest) (*OrganizerEventsResult, error) {
	resp := new(OrganizerEventsResult)

	return resp, c.getJSON(ctx, fmt.Sprintf("/organizers/%s/events/", id), req, resp)
}
