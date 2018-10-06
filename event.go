package eventbrite

import (
	"fmt"
	"net/url"

	"golang.org/x/net/context"
)

// Event is an object representing anything from a small birthday party to a massive stadium
// concert and everything in between.
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-event
type Event struct {
	// Event ID
	Id string `json:"id"`
	// The event’s name
	Name MultipartText `json:"name"`
	// multipart-text: The event’s description (contents of the event page).
	// May be long and have significant formatting. (optional)
	Description MultipartText `json:"description"`
	// The URL to the event page for this event on Eventbrite
	Url string `json:"url"`
	// The start time of the event
	Start DatetimeTz `json:"start"`
	// The end time of the event
	End DatetimeTz `json:"end"`
	// When the event was created
	Created DateTime `json:"created"`
	// When the event was last changed
	Changed DateTime `json:"changed"`
	// One of canceled, live, started, ended, completed
	Status string `json:"status"`
	// The ISO 4217 currency code for this event
	Currency string `json:"currency"`
	// If this event doesn’t have a venue and is only held online
	OnlineEvent bool `json:"online_event"`
	// The venue the event is held at (optional)
	Venue   Venue  `json:"venue"`
	VenueId string `json:"venue_id"`
	// The organizer of the event
	Organizer   Organizer `json:"organizer"`
	OrganizerId string    `json:"organizer_id"`
	// The event’s format (type of event: conference, seminar, concert, etc.) (optional)
	Format   Format `json:"format"`
	FormatId string `json:"format_id"`
	// The event’s category (technology, music, science, etc.) (optional)
	Category   Category `json:"category"`
	CategoryId string   `json:"category_id"`
	// The event’s subcategory (optional)
	SubCategory   SubCategory `json:"subcategory"`
	SubCategoryId string      `json:"subcategory_id"`
	LogoID        string      `json:"logo_id"`
	// The image logo for this event (optional)
	Logo Image `json:"logo"`
	// The event’s refund policy (optional)
	RefundPolicy interface{} `json:"refund_policy"`
	// The bookmark information on the event. Currently returns a dictionary with the number of users who
	// have bookmarked the event as ‘count’ (i.e. {'count': 3})
	BookmarkInfo interface{} `json:"bookmark_info"`
}

// EventSearchRequest is the request structure for searching Event
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-parameters
type EventSearchRequest struct {
	// Return events matching the given keywords. This parameter will accept any string as a keyword.
	Query string `json:"q"`
	// Parameter you want to sort by - options are “date”, “distance” and “best”. Prefix with a
	// hyphen to reverse the order, e.g. “-date”.
	SortBy string `json:"sort_by"`
	// The address of the location you want to search for events around.
	LocationAddress string `json:"location.address"`
	// The distance you want to search around the given location. This should be an integer followed by “mi” or “km”.
	LocationWithin string `json:"location.within"`
	// The latitude of of the location you want to search for events around.
	LocationLatitude string `json:"location.latitude"`
	// The longitude of the location you want to search for events around.
	LocationLongitude string `json:"location.longitude"`
	// The latitude of the northeast corner of a viewport.
	LocationViewportNortheastLatitude string `json:"location.viewport.northeast.latitude"`
	// The longitude of the northeast corner of a viewport.
	LocationViewportNortheastLongitude string `json:"location.viewport.northeast.longitude"`
	// The latitude of the southwest corner of a viewport.
	LocationViewportSouthwestLatitude string `json:"location.viewport.southwest.latitude"`
	// The longitude of the southwest corner of a viewport.
	LocationViewportSouthwestLongitude string `json:"location.viewport.southwest.longitude"`
	// Only return events organized by the given Organizer ID.
	OrganizerId string `json:"organizer.id"`
	// Only return events owned by the given User ID.
	UserId string `json:"user.id"`
	// Append the given tracking_code to the event URLs returned.
	TrackingCode string `json:"tracking_code"`
	// Only return events under the given category IDs. This should be a comma delimited string of category IDs.
	Categories string `json:"categories"`
	// Only return events under the given subcategory IDs. This should be a comma delimited string of subcategory IDs.
	Subcategories string `json:"subcategories"`
	// Only return events with the given format IDs. This should be a comma delimited string of format IDs.
	Formats string `json:"formats"`
	//    Only return events that are “free” or “paid”
	Price string `json:"price"`
	// Only return events with start dates after the given date.
	StartDateRangeStart string `json:"start_date.range_start"`
	// Only return events with start dates before the given date.
	StartDateRangeEnd string `json:"start_date.range_end"`
	// Only return events with start dates within the given keyword date range. Keyword options are “this_week”,
	// “next_week”, “this_weekend”, “next_month”, “this_month”, “tomorrow”, “today”
	StartDateKeyword string `json:"start_date.keyword"`
	// Only return events with modified dates after the given UTC date.
	DateModifiedRangeStart string `json:"date_modified.range_start"`
	// Only return events with modified dates before the given UTC date.
	DateModifiedEnd string `json:"date_modified.range_end"`
	// Only return events with modified dates within the given keyword date range. Keyword options are “this_week”,
	// “next_week”, “this_weekend”, “next_month”, “this_month”, “tomorrow”, “today”
	DateModifiedKeyword string `json:"date_modified.keyword"`
	// Use the preconfigured settings for this type of search - Current option is “promoted”
	SearchType string `json:"search_type"`
	// Boolean for whether or not you want to see all instances of repeating events in search results.
	IncludeAllSeriesInstances bool `json:"include_all_series_instances"`
	// Boolean for whether or not you want to see events without tickets on sale.
	IncludeUnavailableEvents bool `json:"include_unavailable_events"`
	// Boolean for whether or not you want to see adult events
	IncludeAdultEvents bool `json:"include_adult_events"`
	// Incorporate additional information from the user’s historic preferences.
	IncorporateUserAffinities bool `json:"incorporate_user_affinities"`
	// Make search results prefer events in these categories. This should be a comma delimited string of category IDs.
	HighAffinityCategories string `json:"high_affinity_categories"`
}

// EventCreateRequest is the request structure for creating an Event
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id1
type EventCreateRequest struct {
	// The name of the event. Value cannot be empty nor whitespace.
	NameHtml string `json:"event.name.html" validate:"required"`
	// The ID of the organizer of this event
	DescriptionHtml string `json:"event.description.html"`
	// The ID of the organizer of this event
	OrganizerID string `json:"event.organizer_id"`
	// The start time of the event
	StartUtc DateTime `json:"event.start.utc" validate:"required"`
	// Yes Start time timezone (Olson format)
	StartTimezone string `json:"event.start.timezone" validate:"required"`
	// The end time of the event
	EndUtc DateTime `json:"event.end.utc" validate:"required"`
	// End time timezone (Olson format)
	EndTimezone string `json:"event.end.timezone" validate:"required"`
	// Whether the start date should be hidden
	HideStartDate bool `json:"event.hide_start_date"`
	// Whether the end date should be hidden
	HideEndDate bool `json:"event.hide_end_date"`
	// Event currency (3 letter code)
	Currency string `json:"event.currency" validate:"required"`
	// The ID of a previously-created venue to associate with this event. You can omit this field or
	// set it to null if you set online_event.
	VenueId string `json:"event.venue_id"`
	// Is the event online-only (no venue)?
	OnlineEvent bool `json:"event.online_event"`
	// If the event is publicly listed and searchable. Defaults to True.
	Listed bool `json:"event.listed"`
	// The logo for the event
	LogoID string `json:"event.logo_id"`
	// The category (vertical) of the event
	CategoryID string `json:"event.category_id"`
	// The subcategory of the event (US only)
	SubcategoryID string `json:"event.subcategory_id"`
	// The format (general type) of the event
	FormatID string `json:"event.format_id"`
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

// EventUpdateRequest is the request structure for updating an Event
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id5
type EventUpdateRequest struct {
	// The name of the event. Value cannot be empty nor whitespace.
	NameHtml string `json:"event.name.html"`
	// The ID of the organizer of this event
	DescriptionHtml string `json:"event.description.html" validate:"required"`
	// The ID of the organizer of this event
	OrganizerId string `json:"event.organizer_id" validate:"required"`
	// The start time of the event
	StartUtc string `json:"event.start.utc" validate:"required"`
	// Start time timezone (Olson format)
	StartTimezone string `json:"event.start.timezone" validate:"required"`
	// The end time of the event
	EndUtc string `json:"event.end.utc" validate:"required"`
	// End time timezone (Olson format)
	EndTimezone string `json:"event.end.timezone" validate:"required"`
	// Whether the start date should be hidden
	HideStartDate bool `json:"event.hide_start_date"`
	// Whether the end date should be hidden
	HideEndDate bool `json:"event.hide_end_date"`
	// Event currency (3 letter code)
	Currency string `json:"event.currency" validate:"required"`
	// The ID of a previously-created venue to associate with this event. You can omit this field or
	// set it to null if you set online_event.
	VenueID string `json:"event.venue_id"`
	// Is the event online-only (no venue)?
	OnlineEvent bool `json:"event.online_event"`
	// If the event is publicly listed and searchable. Defaults to True.
	Listed bool `json:"event.listed"`
	// The logo for the event
	LogoID string `json:"event.logo_id"`
	// The category (vertical) of the event
	CategoryID string `json:"event.category_id"`
	// The subcategory of the event (US only)
	SubcategoryID string `json:"event.subcategory_id"`
	// The format (general type) of the event
	FormatID string `json:"event.format_id"`
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

// EventUpdateDisplaySettings is the request structure for updating an Event
// display settings
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id18
type EventUpdateDisplaySettings struct {
	// Whether to display the start date on the event listing
	ShowStartDate bool `json:"display_settings.show_start_date"`
	// Whether to display the end date on the event listing
	ShowEndDate bool `json:"display_settings.show_end_date"`
	// Whether to display event start and end time on the event listing
	ShowStartEndTime bool `json:"display_settings.show_start_end_time"`
	// Whether to display the event timezone on the event listing
	ShowTimezone bool `json:"display_settings.show_timezone"`
	// Whether to display a map to the venue on the event listing
	ShowMap bool `json:"display_settings.show_map"`
	// Whether to display the number of remaining tickets
	ShowRemaining bool `json:"display_settings.show_remaining"`
	// Whether to display a link to the organizer’s Facebook profile
	ShowOrganizerFacebook bool `json:"display_settings.show_organizer_facebook"`
	// Whether to display a link to the organizer’s Twitter profile
	ShowOrganizerTwitter bool `json:"display_settings.show_organizer_twitter"`
	// Whether to display which of the user’s Facebook friends are going
	ShowFacebookFriendsGoing bool `json:"display_settings.show_facebook_friends_going"`
	// Which terminology should be used to refer to the event (Valid choices are: tickets_vertical, or endurance_vertical)
	ShowAttendeeList bool `json:"display_settings.show_attendee_list"`
}

// EventGetTicketClass is the request structure to get an Event TicketClass
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id20
type EventGetTicketClass struct {
	// Only return ticket classes valid for the given point of sale (Valid choices are: online, or at_the_door)
	Pos string `json:"pos"`
}

// EventGetTicketClass is the request structure to create an Event TicketClass
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id22
type EventCreateTicketClass struct {
	// Name of this ticket type
	Name string `json:"ticket_class.name"`
	// Description of the ticket
	Description string `json:"ticket_class.description"`
	// Total available number of this ticket
	QuantityTotal int `json:"quantity_total"`
	// Cost of the ticket (currently currency must match event currency) e.g. $45 would be ‘USD,4500’
	Cost Currency `json:"ticket_class.quantity_total"`
	// Is this a donation? (user-supplied cost)
	Donation bool `json:"ticket_class.donation"`
	// If the ticket is a free ticket
	Free bool `json:"ticket_class.free"`
	// Absorb the fee into the displayed cost
	IncludeFee bool `json:"ticket_class.include_fee"`
	// Absorb the payment fee, but show the eventbrite fee
	SplitFee bool `json:"ticket_class.split_fee"`
	// Hide the ticket description on the event page
	HideDescription bool `json:"ticket_class.hide_description"`
	// A list of all supported sales channels ([“online”], [“online”, “atd”], [“atd”])
	SalesChannels []interface{} `json:"ticket_class.sales_channels"`
	// When the ticket is available for sale (leave empty for ‘when event published’)
	SalesStart string `json:"ticket_class.sales_start"`
	// When the ticket stops being on sale (leave empty for ‘one hour before event start’)
	SalesEnd string `json:"ticket_class.sales_end"`
	// The ID of another ticket class - when it sells out, this class will go on sale.
	SalesStartAfter string `json:"ticket_class.sales_start_after"`
	// Minimum number that can be bought per order
	MinimumQuantity int `json:"ticket_class.minimum_quantity"`
	// Maximum number that can be bought per order
	MaximumQuantity int `json:"ticket_class.maximum_quantity"`
	// How many of these tickets have already been sold and confirmed (does not include tickets being checked out right now)
	QuantitySold int `json:"quantity_sold"`
	// Hide this ticket
	Hidden bool `json:"ticket_class.hidden"`
	// Hide this ticket when it is not on sale
	AutoHide bool `json:"ticket_class.auto_hide"`
	// Override reveal date for auto-hide
	AutoHideBefore string `json:"ticket_class.auto_hide_before"`
	// Override re-hide date for auto-hide
	AutoHideAfter string `json:"ticket_class.auto_hide_after"`
	// Order message per ticket type
	OrderConfirmationMessage string `json:"ticket_class.order_confirmation_message"`
}

// EventUpdateTicketClass is the request structure to update an Event TicketClass
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id26
type EventUpdateTicketClass struct {
	// Name of this ticket type
	Name string `json:"ticket_class.name"`
	// Description of the ticket
	Description string `json:"ticket_class.description"`
	// Total available number of this ticket
	QuantityTotal int `json:"quantity_total"`
	// Cost of the ticket (currently currency must match event currency) e.g. $45 would be ‘USD,4500’
	Cost Currency `json:"ticket_class.quantity_total"`
	// Is this a donation? (user-supplied cost)
	Donation bool `json:"ticket_class.donation"`
	// If the ticket is a free ticket
	Free bool `json:"ticket_class.free"`
	// Absorb the fee into the displayed cost
	IncludeFee bool `json:"ticket_class.include_fee"`
	// Absorb the payment fee, but show the eventbrite fee
	SplitFee bool `json:"ticket_class.split_fee"`
	// Hide the ticket description on the event page
	HideDescription bool `json:"ticket_class.hide_description"`
	// A list of all supported sales channels ([“online”], [“online”, “atd”], [“atd”])
	SalesChannels []interface{} `json:"ticket_class.sales_channels"`
	// When the ticket is available for sale (leave empty for ‘when event published’)
	SalesStart string `json:"ticket_class.sales_start"`
	// When the ticket stops being on sale (leave empty for ‘one hour before event start’)
	SalesEnd string `json:"ticket_class.sales_end"`
	// The ID of another ticket class - when it sells out, this class will go on sale.
	SalesStartAfter string `json:"ticket_class.sales_start_after"`
	// Minimum number that can be bought per order
	MinimumQuantity int `json:"ticket_class.minimum_quantity"`
	// Maximum number that can be bought per order
	MaximumQuantity int `json:"ticket_class.maximum_quantity"`
	// How many of these tickets have already been sold and confirmed (does not include tickets being checked out right now)
	QuantitySold int `json:"quantity_sold"`
	// Hide this ticket
	Hidden bool `json:"ticket_class.hidden"`
	// Hide this ticket when it is not on sale
	AutoHide bool `json:"ticket_class.auto_hide"`
	// Override reveal date for auto-hide
	AutoHideBefore string `json:"ticket_class.auto_hide_before"`
	// Override re-hide date for auto-hide
	AutoHideAfter string `json:"ticket_class.auto_hide_after"`
	// Order message per ticket type
	OrderConfirmationMessage string `json:"ticket_class.order_confirmation_message"`
}

// EventDeleteTicketClass is the request structure to delkete an Event TicketClass
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id29
type EventDeleteTicketClass struct {
	// Delete even if ticket sales depend on this ticket. This will start ticket sales of
	// dependents immediately
	BreakDependency bool `json:"break_dependency"`
}

// EventGetCannedQuestions is the request structure to get Event canned questions
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id31
type EventGetCannedQuestions struct {
	// Return private events and more details
	AsOwner bool `json:"as_owner"`
}

// EventCreateCannedQuestion is the request structure to create an Event canned question
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id33
type EventCreateCannedQuestion struct {
	// Question displayed to the recipient
	Html string `json:"question.question.html"`
	// Is an answer to this question required for registration?
	Required bool `json:"question.required"`
	// Type of Question (Valid choices are: checkbox, dropdown, text, paragraph, radio, or waiver)
	Type string `json:"question.type"`
	// Ask this question to the ticket buyer or each attendee? (Valid choices are: ticket_buyer, or attendee)
	Respondent string `json:"question.respondent" validate:"required"`
	// Waiver content for questions of type waiver
	Waiver string `json:"question.waiver"`
	// Choices for multiple choice questions. Format:
	// [{“answer”: {“html”: “Choice goes here...”}}, {“answer”: {“html”: “Another choice goes here...”}}]
	Choices interface{} `json:"question.choices"`
	// Tickets to which to limit this question. Format: [{“id”: “1234”}, {“id”: “4567”}]
	TicketClasses interface{} `json:"question.ticket_classes"`
	// ID of Parent Question (for subquestions)
	ParentChoiceID string `json:"question.parent_choice_id"`
	// Is this question displayed on order confirmation?
	DisplayAnswerOnOrder bool `json:"question.display_answer_on_order"`
	// String value of canned_type
	CannedType string `json:"question.canned_type"`
}

// EventGetQuestions is the request structure to get an Event questions
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id35
type EventGetQuestions struct {
	// Return private events and more details
	AsOwner bool `json:"as_owner"`
}

// EventCreateQuestion is the request structure to create an Event question
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id37
type EventCreateQuestion struct {
	// Question displayed to the recipient
	Html string `json:"question.question.html"`
	// Is an answer to this question required for registration?
	Required bool `json:"question.required"`
	// Type of Question (Valid choices are: checkbox, dropdown, text, paragraph, radio, or waiver)
	Type string `json:"question.type"`
	// Ask this question to the ticket buyer or each attendee? (Valid choices are: ticket_buyer, or attendee)
	Respondent string `json:"question.respondent" validate:"required"`
	// Waiver content for questions of type waiver
	Waiver string `json:"question.waiver"`
	// Choices for multiple choice questions. Format:
	// [{“answer”: {“html”: “Choice goes here...”}}, {“answer”: {“html”: “Another choice goes here...”}}]
	Choices interface{} `json:"question.choices"`
	// Tickets to which to limit this question. Format: [{“id”: “1234”}, {“id”: “4567”}]
	TicketClasses interface{} `json:"question.ticket_classes"`
	// ID of Parent Question (for subquestions)
	ParentChoiceID string `json:"question.parent_choice_id"`
	// Is this question displayed on order confirmation?
	DisplayAnswerOnOrder bool `json:"question.display_answer_on_order"`
}

// EventGetAttendees is the request structure to get an Event Attendees list
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id41
type EventGetAttendees struct {
	// Limits results to either confirmed attendees or cancelled/refunded/etc.
	// attendees (Valid choices are: attending, not_attending, or unpaid)
	Status string `json:"status"`
	// Only return attendees changed on or after the time given
	ChangedSince string `json:"changed_since"`
	// Only return attendees changed on or after the time given and with an id bigger than last item seen
	LastItemSeen int `json:"last_item_seen"`
	// Only return attendees whose ids are in this list
	AttendeeIds []interface{} `json:"attendee_ids"`
}

// EventGetOrders is the request structure to get an Event Order list
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id45
type EventGetOrders struct {
	// Limits results to either confirmed attendees or cancelled/refunded/etc.
	// attendees (Valid choices are: attending, not_attending, or unpaid)
	Status string `json:"status"`
	// Only return attendees changed on or after the time given
	ChangedSince string `json:"changed_since"`
	// Only return attendees changed on or after the time given and with an id bigger than last item seen
	LastItemSeen int `json:"last_item_seen"`
	// Only include orders placed by one of these emails
	OnlyEmails []interface{} `json:"only_emails"`
	// Don’t include orders placed by any of these emails
	ExcludeEmails []interface{} `json:"only_emails"`
	// Return only orders with selected refund requests statuses.
	// Possible values are: completed, pending, outside_policy, disputed, denied
	RefundRequestStatuses []interface{} `json:"refund_request_statuses"`
}

// EventGetTransfers is the request structure to get an Event Transfer list
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-id61
type EventGetTransfers struct {
	ChangedSince string `json:"changed_since"`
}

// EventGetTicketGroups is the request to get an Event TicketGroup list
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-event-id-ticket-groups
type EventGetTicketGroups struct {
	// Limits results to groups with the specific status (Valid choices are: live, archived, deleted, or all)
	Status string `json:"status"`
}

// EventGetTicketGroupsTicketClasses is the request structure to get TicketGroup TicketClass list
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-event-id-ticket-classes-ticket-class-id-ticket-groups
type EventGetTicketGroupsTicketClasses struct {
	// Limits results to groups with the specific status (Valid choices are: live, archived, deleted, or all)
	Status string `json:"status"`
}

// EventSearchResult is the response structure for Event search
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-events
type EventSearchResult struct {
	Pagination     Pagination `json:"pagination"`
	Events         []Event    `json:"events"`
	TopMatchEvents []Event    `json:"top_match_events"`
}

// EventGetTicketClassResult is the response structure for an Event TicketClass
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-ticket-classes
type EventGetTicketClassResult struct {
	Pagination    Pagination    `json:"pagination"`
	TicketClasses []TicketClass `json:"ticket_classes"`
}

// EventSearch allows you to retrieve a paginated response of public event objects from across
// Eventbrite’s directory, regardless of which user owns the event.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-events
func (c *Client) EventSearch(ctx context.Context, req *EventSearchRequest) (*EventSearchResult, error) {
	result := &EventSearchResult{}

	return result, c.getJSON(ctx, "/events/search/", req, &result)
}

// EventGet returns an event for the specified event. Many of Eventbrite’s API use cases revolve around pulling
// details of a specific event within an Eventbrite account. Does not support fetching a repeating event
// series parent (see GET /series/:id/).
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id
func (c *Client) EventGet(ctx context.Context, id string) (*Event, error) {
	result := &Event{}

	return result, c.getJSON(ctx, "/events/"+id, url.Values{}, &result)
}

// EventCreate makes a new event, and returns an event for the specified event. Does not support the
// creation of repeating event series.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events
func (c *Client) EventCreate(ctx context.Context, req *EventCreateRequest) (interface{}, error) {
	var resp interface{}

	return resp, c.postJSON(ctx, "/events/", req, &resp)
}

// EventUpdate updates an event. Returns an event for the specified event. Does not support updating a
// repeating event series parent (see POST /series/:id/)
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id
func (c *Client) EventUpdate(ctx context.Context, id string, req *EventUpdateRequest) (interface{}, error) {
	event := &Event{}

	return event, c.postJSON(ctx, fmt.Sprintf("/events/%s/", id), req, event)
}

// EventPublish publishes an event if it has not already been deleted. In order for publish to be permitted, the event
// must have all necessary information, including a name and description, an organizer, at least one
// ticket, and valid payment options. This API endpoint will return argument errors for event fields that
// fail to validate the publish requirements. Returns a boolean indicating success or failure of the publish.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-events
func (c *Client) EventPublish(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/events/%s/publish", id)

	var resp interface{}
	return resp, c.postJSON(ctx, path, nil, &resp)
}

// EventUnPublish unpublishes an event. In order for a free event to be unpublished, it must not have any pending or completed
// orders, even if the event is in the past. In order for a paid event to be unpublished, it must not have
// any pending or completed orders, unless the event has been completed and paid out. Returns a boolean indicating
// success or failure of the unpublish.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id-unpublish
func (c *Client) EventUnPublish(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/events/%s/unpublish", id)

	var resp interface{}
	return resp, c.postJSON(ctx, path, nil, &resp)
}

// EventCancel cancels an event if it has not already been deleted. In order for cancel to be permitted, there must be no
// pending or completed orders. Returns a boolean indicating success or failure of the cancel.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id-cancel
func (c *Client) EventCancel(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/events/%s/unpublish", id)

	var resp interface{}
	return resp, c.postJSON(ctx, path, nil, &resp)
}

// EventDelete deletes an event if the delete is permitted. In order for a delete to be permitted, there must be no pending
// or completed orders. Returns a boolean indicating success or failure of the delete.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-delete-events-id
func (c *Client) EventDelete(ctx context.Context, id string) (interface{}, error) {
	path := fmt.Sprintf("/events/%s", id)

	var resp interface{}
	return resp, c.deleteJSON(ctx, path, &resp)
}

// EventGetDisplaySettings gets Event display settings
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-display-settings
func (c *Client) EventGetDisplaySettings(ctx context.Context, id string) (*EventSettings, error) {
	result := new(EventSettings)

	return result, c.getJSON(ctx, fmt.Sprintf("/events/%s/display_settings/", id), url.Values{}, &result)
}

// EventUpdateDisplaySettings apdates the display settings for an Event.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id-display-settings
func (c *Client) EventUpdateDisplaySettings(ctx context.Context, id string, settings *EventUpdateDisplaySettings) (*EventSettings, error) {
	result := new(EventSettings)

	return result, c.postJSON(ctx, fmt.Sprintf("/events/%s/display_settings/", id), settings, &result)
}

// EventGetTicketClasses gets an Event TicketClass
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-ticket-classes
func (c *Client) EventGetTicketClasses(ctx context.Context, id string, class *EventGetTicketClass) (*EventGetTicketClassResult, error) {
	result := new(EventGetTicketClassResult)

	return result, c.getJSON(ctx, fmt.Sprintf("/events/%s/ticket_classes/", id), class, result)
}

// EventCreateTicketClass creates a new ticket class, returning the result as a ticket_class under the key ticket_class.
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id-ticket-classes
func (c *Client) EventCreateTicketClass(ctx context.Context, id string, class *EventCreateTicketClass) (*TicketClass, error) {
	result := new(TicketClass)

	return result, c.postJSON(ctx, fmt.Sprintf("/events/%s/ticket_classes/", id), class, result)
}

// EventGetTicketClass gets and returns a single TicketClass by ID
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-ticket-classes-ticket-class-id
func (c *Client) EventGetTicketClass(ctx context.Context, eventId, ticketId string) (*TicketClass, error) {
	result := new(TicketClass)

	return result, c.getJSON(ctx, fmt.Sprintf("/events/%s/ticket_classes/%s/", eventId, ticketId), nil, result)
}

// EventUpdateTicketClass updates an existing ticket class, returning the updated result as a ticket_class under the key
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id-ticket-classes-ticket-class-id
func (c *Client) EventUpdateTicketClass(ctx context.Context, eventId, ticketId string, class *EventUpdateTicketClass) (*TicketClass, error) {
	result := new(TicketClass)

	return result, c.postJSON(ctx, fmt.Sprintf("/events/%s/ticket_classes/%s/", eventId, ticketId), nil, result)
}

// EventDeleteTicketClass deletes the ticket class. Returns {"deleted": true}
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-delete-events-id-ticket-classes-ticket-class-id
func (c *Client) EventDeleteTicketClass(ctx context.Context, eventId, ticketId string, class *EventDeleteTicketClass) (interface{}, error) {
	result := new(TicketClass)

	return result, c.deleteJSON(ctx, fmt.Sprintf("/events/%s/ticket_classes/%s/", eventId, ticketId), result)
}

// EventGetCannedQuestions this endpoint returns canned questions of a single event
// (examples: first name, last name, company, prefix, etc.).
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-canned-questions
func (c *Client) EventGetCannedQuestions(ctx context.Context, id string, q *EventGetCannedQuestions) (interface{}, error) {
	var result interface{}

	return result, c.getJSON(ctx, fmt.Sprintf("/events/%s/canned_questions/", id), q, result)
}

// EventCreateCannedQuestion creates a new canned question; returns the result as a question
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id-canned-questions
func (c *Client) EventCreateCannedQuestion(ctx context.Context, id string, q *EventCreateCannedQuestion) (interface{}, error) {
	var result interface{}

	return result, c.postJSON(ctx, fmt.Sprintf("/events/%s/canned_questions/", id), q, result)
}

// Eventbrite allows event organizers to add custom questions that attendees fill out upon registration.
// This endpoint can be helpful for determining what custom information is collected and available per event.
// This endpoint will return question
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-questions
func (c *Client) EventGetQuestions(ctx context.Context, id string, q *EventGetQuestions) (interface{}, error) {
	var result interface{}

	return result, c.postJSON(ctx, fmt.Sprintf("/events/%s/questions/", id), q, result)
}

// EventCreateQuestion creates a new question; returns the result as a question as the key question
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-post-events-id-questions
func (c *Client) EventCreateQuestion(ctx context.Context, id string, q *EventCreateQuestion) (interface{}, error) {
	var result interface{}

	return result, c.postJSON(ctx, fmt.Sprintf("/events/%s/questions/", id), q, result)
}

// EventGetQuestion returns question for a specific question id
//
// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-questions-id
func (c *Client) EventGetQuestion(ctx context.Context, eventId, questionId string) (interface{}, error) {
	var result interface{}

	return result, c.postJSON(ctx, fmt.Sprintf("/events/%s/questions/%s/", eventId, questionId), nil, result)
}
