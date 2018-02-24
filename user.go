package eventbrite

import (
	"context"
	"fmt"
)

// User is an object representing an Eventbrite user
//
// https://www.eventbrite.com/developer/v3/response_formats/user/#ebapi-std:format-user
type User struct {
	ID string `json:"id"`
	// The user’s name. Use this in preference to first_name/last_name if possible for forward compatibility with non-Western names
	Name string `json:"name"`
	// The user’s first name
	FirstName string `json:"first_name"`
	// The user’s last name
	LastName string `json:"last_name"`
	// A list of user emails
	Emails []Email `json:"emails"`
}

// Email contains a list of email objects giving information on the user’s email addresses
//
// https://www.eventbrite.com/developer/v3/response_formats/user/#ebapi-user-emails
type Email struct {
	Email string `json:"email"`
	Verified bool `json:"verified"`
	Primary bool `json:"primary"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id1
type GetUserOrdersRequest struct {
	// Only return resource changed on or after the time given
	ChangedSince string `json:"changed_since"`
	// Limits results to either past or current & future events / orders.
	// (Valid choices are: all, past, or current_future)
	TimeFilter string `json:"time_filter"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id3
type GetUserOrganizersRequest struct {
	// 	True: Will hide organizers flagged as “unsaved” False: Will show organizers
	// regardless of unsaved flag (Default value)
	HideUnsaved bool `json:"hide_unsaved"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id5
type GetUserOwnedEvents struct {
	// How to order the results (Valid choices are: start_asc, start_desc, created_asc,
	// created_desc, name_asc, or name_desc)
	OrderBy string `json:"order_by"`
	// True: Will show parent of a serie instead of children False: Will show children of a serie (Default value)
	ShowSeriesParent bool `json:"show_series_parent"`
	// Filter by events with a specific status set. This should be a comma delimited string of status.
	// Valid status: all, draft, live, canceled, started, ended.
	Status string `json:"status"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id7
type GetUserEvents struct {
	// Filter event results by name
	NameFilter string `json:"name_filter"`
    // Filter event results by currency
	CurrencyFilter string `json:"currency_filter"`
    // How to order the results (Valid choices are: start_asc, start_desc, created_asc,
	// created_desc, name_asc, or name_desc)
	OrderBy string `json:"order_by"`
    // True: Will show parent of a serie instead of children False: Will show children of a serie (Default value)
	ShowSeriesParent bool `json:"show_series_parent"`
    // Filter by events with a specific status set. This should be a comma delimited string of status.
	// Valid status: all, draft, live, canceled, started, ended
	Status string `json:"status"`
    // Filter event results by event_group_id
	EventGroupID string `json:"event_group_id"`
    // Number of records in each page
	PageSize int `json:"page_size"`
    // Limits results to either past or current & future events / orders. (Valid choices are: all, past, or current_future
	TimeFilter string `json:"time_filter"`
    // Filter event results by venue IDs
	VenueFilter []interface{} `json:"venue_filter"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id9
type CreateOrganizationEventRequest struct {
	// The name of the event. Value cannot be empty nor whitespace.
	NameHtml string `json:"event.name.html" validate:"required"`
    // The ID of the organizer of this event
	DescriptionHtml string	`json:"event.description.html" validate:"required"`
    // The ID of the organizer of this event
	OrganizerId string `json:"event.organizer_id" validate:"required"`
    // The start time of the event
	StartUtc string `json:"event.start.utc" validate:"required"`
    // Yes	Start time timezone (Olson format)
	EventStartTimezone string `json:"event.start.timezone" validate:"required"`
    // The end time of the event
	EventEndUtc string `json:"event.end.utc" validate:"required"`
    //	End time timezone (Olson format)
	EventEndTimezone string `json:"event.end.timezone" validate:"required"`
    // Whether the start date should be hidden
	EventHideStartDate bool `json:"event.hide_start_date"`
    // Whether the end date should be hidden
	EventHideEndDate bool `json:"event.hide_end_date"`
    // Event currency (3 letter code)
	EventCurrency string `json:"event.currency" validate:"required"`
    // The ID of a previously-created venue to associate with this event. You can omit this field or
	// set it to null if you set online_event.
	VenueId string `json:"event.venue_id"`
    // Is the event online-only (no venue)?
	OnlineEvent bool `json:"event.online_event"`
    // If the event is publicly listed and searchable. Defaults to True.
	Listed bool `json:"event.listed"`
    // The logo for the event
	LogoId string `json:"event.logo_id"`
    // The category (vertical) of the event
	CategoryId string `json:"event.category_id"`
    // The subcategory of the event (US only)
	SubcategoryId string `json:"event.subcategory_id"`
    // The format (general type) of the event
	FormatId string `json:"event.format_id"`
    // If users can share the event on social media
	Sharable bool `json:"event.shareable"`
    // Only invited users can see the event page
	InviteOnly bool `json:"event.invite_only"`
    // Password needed to see the event in unlisted mode
	Password string `json:"event.password"`
    // Set specific capacity (if omitted, sums ticket capacities)
	Capacity int `json:"event.capacity"`
    // If the remaining number of tickets is publicly visible on the event page
	ShowRemaining bool `json:"event.show_remaining"`
    // If the event is reserved seating
	IsReservedSeating bool `json:"event.is_reserved_seating"`
    // Source of the event (defaults to API)
	Source string `json:"event.source"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-venues
type GetUserVenuesResult struct {
	Pagination Pagination `json:"pagination"`
	Venues []Venue `json:"venues"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id13
type CreateOrganizationVenueRequest struct {
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

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id15
type UserEventAttendees struct {
	// Limits results to either confirmed attendees or cancelled/refunded/etc. attendees
	// (Valid choices are: attending, or not_attending)
	Status string `json:"status"`
    // Only return resource changed on or after the time given
	ChangedSince string `json:"changed_since"`
}

// UserEventOrders is the request structure to get all order placed under
// the user
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id17
type UserEventOrders struct {
	// Limits results to either past or current & future events / orders.
	// (Valid choices are: all, past, or current_future)
	TimeFilter string `json:"time_filter"`
	// Only return resource changed on or after the time given
	ChangedSince string `json:"changed_since"`
}

// GetUserOrdersResult is the response structure for user orders
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-orders
type UserOrdersResult struct {
	Pagination Pagination `json:"pagination"`
	Orders []Order `json:"orders"`
}

// UserOrganizerRequest is the request structure to get all organizer objects that are owned by the user
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-id3
type UserOrganizerRequest struct {
	// True: Will hide organizers flagged as “unsaved” False: Will show organizers
	// regardless of unsaved flag (Default value)
	HideUnsaved bool `json:"hide_unsaved"`
}

// UserOrganizerResponse is the response structure for all organizer objects that are owned by the user
type UserOrganizerResponse struct {
	Pagination Pagination `json:"pagination"`
	Organizers []Organizer `json:"organizers"`
}

// UserOwnedEventsRequest is the request structure to get user owned events
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-id5
type UserOwnedEventsRequest struct {
	// How to order the results (Valid choices are: start_asc, start_desc, created_asc,
	// created_desc, name_asc, or name_desc)
	OrderBy string `json:"order_by"`
	// True: Will show parent of a serie instead of children False: Will show children of a serie (Default value)
	ShowSeriesParent bool `json:"show_series_parent"`
	// Filter by events with a specific status set. This should be a comma delimited string of status.
	// Valid status: all, draft, live, canceled, started, ended
	Status string `json:"status"`
}

// UserOwnedEventResponse is the response structure to get user owned events
type UserOwnedEventResponse struct {
	Pagination Pagination `json:"pagination"`
	Events []Event `json:"events"`
}

type UserEventsRequest struct {

}

type UserEventsResponse struct {
	// Filter event results by name
	NameFilter string `json:"name_filter"`
	// Filter event results by currency
	CurrencyFilter string `json:"currency_filter"`
	// How to order the results (Valid choices are: start_asc, start_desc, created_asc, created_desc, name_asc, or name_desc)
	OrderBy string `json:"order_by"`
	// True: Will show parent of a serie instead of children False: Will show children of a serie (Default value)
	ShowSeriesParent bool `json:"show_series_parent"`
	// Filter by events with a specific status set. This should be a comma delimited string of
	// status. Valid status: all, draft, live, canceled, started, ended.
	Status string `json:"status"`
	// Filter event results by event_group_id
	EventGroupID string `json:"event_group_id"`
	// Number of records in each page.
	PageSize int `json:"page_size"`
	// Limits results to either past or current & future events / orders. (Valid choices are: all, past, or current_future)
	TimeFilter string `json:"time_filter"`
	// Filter event results by venue IDs
	VenueFilter []interface{} `json:"venue_filter"`
}

// UserVenuesResponse is the reponse structure to get user owned venues
type UserVenuesResponse struct {
	Pagination Pagination `json:"pagination"`
	Venues []Venue `json:"venues"`
}

// UserGet returns a user for the specified user as user. If you want to get details about the
// currently authenticated user, use /users/me/
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id
func (c *Client) User(ctx context.Context, id string) (*User, error) {
	u := new (User)

	return u, c.getJSON(ctx, fmt.Sprintf("/users/%s/", id), nil, u)
}

// UserOrders returns a paginated response of orders, under the key orders, of all orders
// the user has placed (i.e. where the user was the person buying the tickets).
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-orders
func (c *Client) UserOrders(ctx context.Context, id string, req *UserEventOrders) (*UserOrdersResult, error) {
	r := new (UserOrdersResult)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/orders/", id), req, r)
}

// UserOrganizers returns a paginated response of organizer objects that are owned by the user.
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-organizers
func (c *Client) UserOrganizers(ctx context.Context, id string, req *UserOrganizerRequest) (*UserOrganizerResponse, error) {
	r := new (UserOrganizerResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/organizers/", id), req, r)
}

// UserOrganizers returns a paginated response of organizer objects that are owned by the user.
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-organizers
func (c *Client) UserOwnedEvents(ctx context.Context, id string, req *UserOwnedEventsRequest) (*UserOwnedEventResponse, error) {
	r := new (UserOwnedEventResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/owned_events/", id), req, r)
}

// UserEvents returns a paginated response of events, under the key events, of all events the user has access to
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-events
func (c *Client) UserEvents(ctx context.Context, id string, req UserEventsRequest) (*UserEventsResponse, error) {
	r := new (UserEventsResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/events/", id), req, r)
}

// UserVenues returns a paginated response of venue objects that are owned by the user
func (c *Client) UserVenues(ctx context.Context, id string) (*UserVenuesResponse, error) {
	r := new (UserVenuesResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/events/", id), nil, r)
}


