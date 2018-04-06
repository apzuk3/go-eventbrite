package eventbrite

import (
	"fmt"

	"golang.org/x/net/context"
)

// https://www.eventbrite.com/developer/v3/endpoints/venues/#ebapi-id1
type UpdateVenueRequest struct {
	// The name of the venue
	Name string `json:"venue.name"`
	// The organizer this venue belongs to (optional - leave this off to use the default organizer)
	OrganizerID string `json:"venue.organizer_id"`
	// The first line of the address
	Address1 string `json:"venue.address.address_1"`
	// The second line of the address
	Address2 string `json:"venue.address.address_2"`
	// The city where the venue is
	City string `json:"venue.address.city"`
	// The region where the venue is
	Region string `json:"venue.address.region"`
	// The postal_code where the venue is
	PostalCode string `json:"venue.address.postal_code"`
	// The country where the venue is
	Country string `json:"venue.address.country"`
	// The latitude of the coordinates for the venue
	Latitude float64 `json:"venue.address.latitude"`
	// The longitude of the coordinates for the venue
	Longitude float64 `json:"venue.address.longitude"`
	// The age restrictions for the venue
	AgeRestriction string `json:"venue.age_restriction"`
	// The max capacity for the venue
	Capacity int `json:"venue.capacity"`
}

// https://www.eventbrite.com/developer/v3/endpoints/venues/#ebapi-id3
type CreateVenueRequest struct {
	// The name of the venue
	Name string `json:"venue.name" validate:"required"`
	// The organizer this venue belongs to (optional - leave this off to use the default organizer)
	OrganizerID string `json:"venue.organizer_id"`
	// The first line of the address
	Address1 string `json:"venue.address.address_1"`
	// The second line of the address
	Address2 string `json:"venue.address.address_2"`
	// The city where the venue is
	City string `json:"venue.address.city"`
	// The region where the venue is
	Region string `json:"venue.address.region"`
	// The postal_code where the venue is
	PostalCode string `json:"venue.address.postal_code"`
	// The country where the venue is
	Country string `json:"venue.address.country"`
	// The latitude of the coordinates for the venue
	Latitude float64 `json:"venue.address.latitude"`
	// The longitude of the coordinates for the venue
	Longitude float64 `json:"venue.address.longitude"`
	// The age restrictions for the venue
	AgeRestriction string `json:"venue.age_restriction"`
	// The max capacity for the venue
	Capacity int `json:"venue.capacity"`
}

type VenueEventsResult struct {
	Pagination Pagination `json:"pagination"`
	Events     []Event    `json:"events"`
}

// https://www.eventbrite.com/developer/v3/endpoints/venues/#ebapi-id5
type GetVenueEventsRequest struct {
	Status              string `json:"status"`
	OrderBy             string `json:"order_by"`
	StartDateRangeStart string `json:"start_date.range_start"`
	StartDateRangeEnd   string `json:"start_date.range_end"`
	OnlyPublic          bool   `json:"only_public"`
}

// Returns a venue object
//
// https://www.eventbrite.com/developer/v3/endpoints/venues/#ebapi-get-venues-id
func (c *Client) VenueGet(ctx context.Context, id string) (*Venue, error) {
	res := new(Venue)

	return res, c.getJSON(ctx, fmt.Sprintf("/venues/%s/", id), nil, res)
}

// Updates a venue and returns it as an object
//
// https://www.eventbrite.com/developer/v3/endpoints/venues/#ebapi-post-venues-id
func (c *Client) VenueUpdate(ctx context.Context, id string, req *UpdateVenueRequest) (*Venue, error) {
	res := new(Venue)

	return res, c.getJSON(ctx, fmt.Sprintf("/venues/%s/", id), nil, res)
}

// Creates a new venue with associated address
//
// https://www.eventbrite.com/developer/v3/endpoints/venues/#ebapi-post-venues
func (c *Client) VenueCreate(ctx context.Context, req *CreateVenueRequest) (*Venue, error) {
	res := new(Venue)

	return res, c.postJSON(ctx, "/venues/", nil, res)
}

// Creates a new venue with associated address
//
// https://www.eventbrite.com/developer/v3/endpoints/venues/#ebapi-post-venues
func (c *Client) VenueEvents(ctx context.Context, venueId string) (*VenueEventsResult, error) {
	res := new(VenueEventsResult)

	return res, c.postJSON(ctx, fmt.Sprintf("/venues/%s/events/", venueId), nil, res)
}
