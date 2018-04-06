package eventbrite

import (
	"fmt"
	"golang.org/x/net/context"
)

// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-parameters
type SeriesCreateEventRequest struct {
	// The name of the event
	Name string `json:"series_parent.name.html" validate:"required"`
	// The description on the event page
	Description string `json:"series_parent.description.html"`
	// of the organizer of this event
	OrganizerID string `json:"series_parent.organizer_id"`
	// The start time of the event
	StartUtc DateTime `json:"series_parent.start.utc" validate:"required"`
	// Start time timezone (Olson format)
	StartTimezone string `json:"series_parent.start.timezone" validate:"required"`
	// The end time of the event
	EndUtc DateTime `json:"series_parent.end.utc" validate:"required"`
	// End time timezone (Olson format)
	EndTimezone string `json:"series_parent.end.timezone" validate:"required"`
	// Whether the start date should be hidden
	HideStartDate bool `json:"series_parent.hide_start_date"`
	// Whether the end date should be hidden
	HideEndDate bool `json:"series_parent.hide_end_date"`
	// Event currency (3 letter code)
	Currency string `json:"series_parent.currency" validate:"required"`
	// ID of the venue
	VenueID string `json:"series_parent.venue_id"`
	// Is the event online-only (no venue)?
	OnlineEvent bool `json:"series_parent.online_event"`
	// If the event is publicly listed and searchable. Defaults to true
	Listed bool `json:"series_parent.listed"`
	// (Deprecated) The logo for the event
	LogoID string `json:"series_parent.logo.id"`
	// The category (vertical) of the event
	CategoryID string `json:"series_parent.category_id"`
	// The subcategory of the event (US only)
	SubCategoryID string `json:"series_parent.subcategory_id"`
	// The format (general type) of the event
	FormatID string `json:"series_parent.format_id"`
	// If users can share the event on social media
	Sharable bool `json:"series_parent.shareable"`
	// Only invited users can see the event page
	InviteOnly bool `json:"series_parent.invite_only"`
	// Password needed to see the event in unlisted mode
	Password string `json:"series_parent.password"`
	// Set specific capacity (if omitted, sums ticket capacities)
	Capacity int `json:"series_parent.capacity"`
	// If the remaining number of tickets is publicly visible on the event page
	ShowRemaining bool `json:"series_parent.show_remaining"`
	// A list of dates for which child events should be created. In the format:
	// [
	//   {
	//     "start": { "utc": "2015-06-15T12:00:00Z", "timezone": "America/Los_Angeles" },
	//     "end": { "utc": "2015-06-15T13:00:00Z", "timezone": "America/Los_Angeles" } },
	//     { ... },
	//     ...
	// ]
	//
	// https://www.eventbrite.co.uk/developer/v3/response_formats/basic/#ebapi-std:format-objectlist
	CreateChildren interface{} `json:"create_children" validate:"required"`
}

type ObjectList []interface{}

// SeriesEventRequest is the response structure for series event
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-id14
type SeriesEventRequest struct {
	// Limits results to either past or current & future events. (Valid choices are: all, past, or current_future)
	TimeFilter string `json:"time_filter"`
	// Append the given tracking_code to the event URLs returned
	TrackingCode string `json:"tracking_code"`
	// How to order the results (Valid choices are: start_asc, start_desc, created_asc, or created_desc)
	OrderBy string `json:"order_by"`
}

// SeriesCUREventRequest is the request structure to make create, update, delete
// requests
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-id16
type SeriesCUREventRequest struct {
	// A list of dates for which child events should be created. In the format:
	//
	//  [
	//    {
	//      "start": { "utc": "2015-06-15T12:00:00Z", "timezone": "America/Los_Angeles" },
	//      "end": { "utc": "2015-06-15T13:00:00Z", "timezone": "America/Los_Angeles" }
	//    },
	//    { ... },
	//    ...
	//  ]
	//
	// https://www.eventbrite.co.uk/developer/v3/response_formats/basic/#ebapi-std:format-objectlist
	CreateChildren interface{} `json:"create_children"`
	// A map of event IDs to modified date objects for updating child events. In the format:
	//
	// {
	//   "1234": { "start": { "utc": "2015-06-15T12:00:00Z", "timezone": "America/Los_Angeles" },
	//   "end": { "utc": "2015-06-15T13:00:00Z", "timezone": "America/Los_Angeles" } },
	//   "5678": { ... },
	//   ...
	// }
	UpdateChildren interface{} `json:"create_children"`
	// A list of IDs for child events that should be deleted. In the format: 1234,5678,9012
	DeleteChildren []string `json:"delete_children"`
}

// EventSeriesCreate creates a new repeating event series. The POST data must include information for at
// least one event date in the series.
//
// Return object is not documented
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-post-series
func (c *Client) EventSeriesCreate(ctx context.Context, req *SeriesCreateEventRequest) (interface{}, error) {
	var resp interface{}

	return resp, c.postJSON(ctx, "/series/", req, &resp)
}

// EventSeriesGet returns a repeating event series parent object for the specified repeating event series
//
// Return object is not documented
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-get-series-id
func (c *Client) EventSeriesGet(ctx context.Context, id string) (interface{}, error) {
	var resp interface{}

	return resp, c.getJSON(ctx, fmt.Sprintf("/series/%s", id), nil, &resp)
}

// Publishes a repeating event series and all of its occurrences that are not already canceled or deleted.
// Once a date is cancelled it can still be uncancelled and can be viewed by the public. A deleted date
// cannot be undeleted and cannot by viewed by the public. In order for publish to be permitted, the event
// must have all necessary information, including a name and description, an organizer, at least one ticket,
// and valid payment options. This API endpoint will return argument errors for event fields that fail to
// validate the publish requirements. Returns a boolean indicating success or failure of the publish
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-post-series-id-publish
func (c *Client) EventSeriesPublish(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/series/%s/publish", id)

	var resp interface{}
	return resp, c.postJSON(ctx, path, nil, &resp)
}

// Unpublishes a repeating event series and all of its occurrences that are not already completed, canceled,
// or deleted. In order for a free series to be unpublished, it must not have any pending or completed orders
// for any dates, even past dates. In order for a paid series to be unpublished, it must not have any pending
// or completed orders for any dates, except that completed orders for past dates that have been completed and
// paid out do not prevent an unpublish. Returns a boolean indicating success or failure of the unpublish
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-post-series-id-unpublish
func (c *Client) EventSeriesUnPublish(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/series/%s/unpublish", id)

	var resp interface{}
	return resp, c.postJSON(ctx, path, nil, &resp)
}

// Cancels a repeating event series and all of its occurrences that are not already canceled or deleted. In order
// for cancel to be permitted, there must be no pending or completed orders for any dates in the series. Returns
// a boolean indicating success or failure of the cancel
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-post-series-id-cancel
func (c *Client) EventSeriesCancel(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/series/%s/unpublish", id)

	var resp interface{}
	return resp, c.postJSON(ctx, path, nil, &resp)
}

// Deletes a repeating event series and all of its occurrences if the delete is permitted. In order for a delete to
// be permitted, there must be no pending or completed orders for any dates in the series. Returns a boolean
// indicating success or failure of the delete
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-delete-series-id
func (c *Client) EventSeriesDelete(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/series/%s", id)

	var resp interface{}
	return resp, c.deleteJSON(ctx, path, &resp)
}

// Creates more event dates or updates or deletes existing event dates in a repeating event series. In order for a
// series date to be deleted or updated, there must be no pending or completed orders for that date
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/events_series/#ebapi-post-series-id-events
func (c *Client) EventSeriesCUD(ctx context.Context, id string, req *SeriesCUREventRequest) (interface{}, error) {
	var v interface{}

	return v, c.postJSON(ctx, fmt.Sprintf("/series/%s/events/", id), req, v)
}
