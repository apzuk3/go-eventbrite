package eventbrite

import (
	"fmt"
	"golang.org/x/net/context"
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

type Contact struct {
	// The contact’s name. Use this in preference to first_name/last_name if possible for
	// forward compatability with non-Western names
	Name string `json:"name"`
	// The contact’s first name
	FirstName string `json:"first_name"`
	// The contact’s last name
	LastName string `json:"last_name"`
	// The contact’s email address
	Email string `json:"email"`
	// When this contact was created
	Created DateTime `json:"created"`
}

// https://www.eventbrite.com/developer/v3/response_formats/user/#ebapi-contact-list
type ContactList struct {
	// The name of the contact list
	Name string `json:"name"`
	// The user who owns this contact list
	UserID string `json:"user_id"`
}

// Email contains a list of email objects giving information on the user’s email addresses
//
// https://www.eventbrite.com/developer/v3/response_formats/user/#ebapi-user-emails
type Email struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	Primary  bool   `json:"primary"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id1
type GetUserOrdersRequest struct {
	// Only return resource changed on or after the time given
	ChangedSince string `json:"changed_since"`
	// Limits results to either past or current & future events / orders.
	// (Valid choices are: all, past, or current_future)
	TimeFilter string `json:"time_filter"`
}

// An assortment is a package/pricing plan associated with an Eventbrite organizer.
// This plan determines the features available to the organizer and the pricing model
// applied to their event tickets.
//
// https://www.eventbrite.com/developer/v3/response_formats/assortments/#ebapi-fields
type Assortment struct {
	// The assortment plan associated with this user
	Plan string `json:"plan"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id3
type GetUserOrganizersRequest struct {
	//     True: Will hide organizers flagged as “unsaved” False: Will show organizers
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
	DescriptionHtml string `json:"event.description.html" validate:"required"`
	// The ID of the organizer of this event
	OrganizerId string `json:"event.organizer_id" validate:"required"`
	// The start time of the event
	StartUtc string `json:"event.start.utc" validate:"required"`
	// Yes    Start time timezone (Olson format)
	EventStartTimezone string `json:"event.start.timezone" validate:"required"`
	// The end time of the event
	EventEndUtc string `json:"event.end.utc" validate:"required"`
	//    End time timezone (Olson format)
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
	Venues     []Venue    `json:"venues"`
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

// UserEventAttendeesRequest is the request structure to get a user owned event attendees
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id15
type UserEventAttendeesRequest struct {
	// Limits results to either confirmed attendees or cancelled/refunded/etc. attendees
	// (Valid choices are: attending, or not_attending)
	Status string `json:"status"`
	// Only return resource changed on or after the time given
	ChangedSince string `json:"changed_since"`
}

// UserEventAttendeesResponse is the response structure to get a user owned event attendees
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-owned-event-attendees
type UserEventAttendeesResponse struct {
	Pagination Pagination `json:"pagination"`
	Attendees  []Attendee `json:"attendee"`
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
	Orders     []Order    `json:"orders"`
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
	Pagination Pagination  `json:"pagination"`
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
	Events     []Event    `json:"events"`
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

// UserVenuesResponse is the response structure to get user owned venues
type UserVenuesResponse struct {
	Pagination Pagination `json:"pagination"`
	Venues     []Venue    `json:"venues"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id17
type UserEventOrdersRequest struct {
	Status        string   `json:"status"`
	OnlyEmails    []string `json:"only_emails"`
	ExcludeEmails []string `json:"exclude_emails"`
	ChangedSince  DateTime `json:"datetime"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-owned-event-orders
type UserEventOrdersResponse struct {
	Pagination Pagination `json:"pagination"`
	Orders     []Order    `json:"orders"`
}

// UserContactListsResponse is the response structure to get user contact lists
type UserContactListsResponse struct {
	Pagination  Pagination    `json:"pagination"`
	ContactList []ContactList `json:"contact_lists"`
}

type UserCreateContactListsRequest struct {
	Name string `json:"contact_list.name" validate:"required"`
}

type UserUpdateContactListRequest struct {
	Name string `json:"contact_list.name" validate:"required"`
}

type UserAddContactListContactRequest struct {
	// Contact’s email address
	Email string `json:"email" validate:"required"`
	// Contact’s first name (or full name)
	FirstName string `json:"first_name"`
	// Contact’s last name
	LastName string `json:"last_name"`
}

type UserDeleteContactListContactRequest struct {
	// Email address to remove
	Email string `json:"email"`
}

type UserContactListContacts struct {
	Pagination Pagination `json:"pagination"`
	Contacts   []Contact  `json:"contacts"`
}

// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-id35
type UserBookmarksRequest struct {
	// Optional bookmark list id to fetch all bookmarks from
	BookmarkListID string `json:"bookmark_list_id"`
}

type UserBookmarksResponse struct {
	Pagination Pagination `json:"pagination"`
	Events     []Event    `json:"events"`
}

type UserSaveBookmarkRequest struct {
	// Event id to bookmark for the user
	EventID int `json:"event_id"`
	// Event ids to batch bookmark for the user
	EventIDs []string `json:"event_ids"`
	// Optional Bookmark list id to save the bookmark(s) to
	BookmarkListID string `json:"bookmark_list_id"`
}

type UserUnSaveBookmarkRequest struct {
	// Event id to bookmark for the user
	EventID int `json:"event_id"`
	// Event ids to batch bookmark for the user
	EventIDs []string `json:"event_ids"`
	// Optional Bookmark list id to save the bookmark(s) to
	BookmarkListID string `json:"bookmark_list_id"`
}

type UserTicketGroupsRequest struct {
	//     Limits results to groups with the specific status (Valid choices are: live, archived, deleted, or all)
	Status string `json:"status"`
}

type UserTicketGroupResponse struct {
	Pagination   Pagination `json:"pagination"`
	TicketGroups []*TicketGroup
}

type UserSetAssortmentRequest struct {
	// The assortments package to upgrade/downgrade to. (Valid choices are: package1, or package2)
	Plan string `json:"plan" validate:"required"`
}

// UserGet returns a user for the specified user as user. If you want to get details about the
// currently authenticated user, use /users/me/
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id
func (c *Client) User(ctx context.Context, id string) (*User, error) {
	u := new(User)

	return u, c.getJSON(ctx, fmt.Sprintf("/users/%s/", id), nil, u)
}

// UserOrders returns a paginated response of orders, under the key orders, of all orders
// the user has placed (i.e. where the user was the person buying the tickets).
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-orders
func (c *Client) UserOrders(ctx context.Context, id string, req *UserEventOrders) (*UserOrdersResult, error) {
	r := new(UserOrdersResult)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/orders/", id), req, r)
}

// UserOrganizers returns a paginated response of organizer objects that are owned by the user.
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-organizers
func (c *Client) UserOrganizers(ctx context.Context, id string, req *UserOrganizerRequest) (*UserOrganizerResponse, error) {
	r := new(UserOrganizerResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/organizers/", id), req, r)
}

// UserOrganizers returns a paginated response of organizer objects that are owned by the user.
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-organizers
func (c *Client) UserOwnedEvents(ctx context.Context, id string, req *UserOwnedEventsRequest) (*UserOwnedEventResponse, error) {
	r := new(UserOwnedEventResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/owned_events/", id), req, r)
}

// UserEvents returns a paginated response of events, under the key events, of all events the user has access to
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/users/#ebapi-get-users-id-events
func (c *Client) UserEvents(ctx context.Context, id string, req UserEventsRequest) (*UserEventsResponse, error) {
	r := new(UserEventsResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/events/", id), req, r)
}

// UserVenues returns a paginated response of venue objects that are owned by the user
func (c *Client) UserVenues(ctx context.Context, id string) (*UserVenuesResponse, error) {
	r := new(UserVenuesResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/events/", id), nil, r)
}

// UserEventAttendees returns a paginated response of attendees, under the key attendees, of attendees visiting
// any of the events the user owns (events that would be returned from /users/:id/owned_events/)
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-owned-event-attendees
func (c *Client) UserEventAttendees(ctx context.Context, id string, request *UserEventAttendeesRequest) (*UserEventAttendeesResponse, error) {
	r := new(UserEventAttendeesResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/owned_event_attendees/", id), request, r)

}

// UserEventOrders returns a paginated response of orders, under the key orders, of orders placed against any of
// the events the user owns (events that would be returned from /users/:id/owned_events/)
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-owned-event-orders
func (c *Client) UserEventOrders(ctx context.Context, id string, request *UserEventOrdersRequest) (*UserEventOrdersResponse, error) {
	r := new(UserEventOrdersResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/owned_event_orders/", id), request, r)

}

// UserContactLists returns a list of contact_list that the user owns as the key contact_lists
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-contact-lists
func (c *Client) UserContactLists(ctx context.Context, id string) (*UserContactListsResponse, error) {
	r := new(UserContactListsResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/contact_lists/", id), nil, r)
}

// UserCreateContactList makes a new contact_list for the user and returns it as contact_list
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-post-users-id-contact-lists
func (c *Client) UserCreateContactList(ctx context.Context, id string, request *UserCreateContactListsRequest) (*UserContactListsResponse, error) {
	r := new(UserContactListsResponse)

	return r, c.postJSON(ctx, fmt.Sprintf("/users/%s/owned_event_orders/", id), request, r)
}

// UserContactList gets a user’s contact_list by ID as contact_list
//
// hhttps://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-contact-lists-contact-list-id
func (c *Client) UserContactList(ctx context.Context, id, contactListID string, request *UserCreateContactListsRequest) (*UserContactListsResponse, error) {
	r := new(UserContactListsResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/contact_lists/%s/", id, contactListID), request, r)
}

// UserUpdateContactList updates the contact_list and returns it as contact_list
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-post-users-id-contact-lists-contact-list-id
func (c *Client) UserUpdateContactList(ctx context.Context, id, contactListID string, request *UserUpdateContactListRequest) (*UserContactListsResponse, error) {
	r := new(UserContactListsResponse)

	return r, c.postJSON(ctx, fmt.Sprintf("/users/%s/contact_lists/%s/", id, contactListID), request, r)
}

// UserDeleteContactList deletes the contact list. Returns {"deleted": true}
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-delete-users-id-contact-lists-contact-list-id
func (c *Client) UserDeleteContactList(ctx context.Context, id, contactListID string) (interface{}, error) {
	var r interface{}

	return r, c.deleteJSON(ctx, fmt.Sprintf("/users/%s/contact_lists/%s/", id, contactListID), r)
}

// UserContactListContacts returns the contacts on the contact list as contacts
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-contact-lists-contact-list-id-contacts
func (c *Client) UserListContactContacts(ctx context.Context, id, contactListID string) (*UserContactListContacts, error) {
	r := new(UserContactListContacts)

	return r, c.postJSON(ctx, fmt.Sprintf("/users/%s/contact_lists/%s/contacts/", id, contactListID), nil, r)
}

// UserContactListContacts adds a new contact to the contact list. Returns {"created": true}
// There is no way to update entries in the list; just delete the old one and add the updated version.
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-contact-lists-contact-list-id-contacts
func (c *Client) UserListContactAddContacts(ctx context.Context, id, contactListID string, req *UserAddContactListContactRequest) (*UserContactListContacts, error) {
	r := new(UserContactListContacts)

	return r, c.postJSON(ctx, fmt.Sprintf("/users/%s/contact_lists/%s/contacts/", id, contactListID), req, r)
}

// UserContactListContacts adds a new contact to the contact list. Returns {"created": true}
// There is no way to update entries in the list; just delete the old one and add the updated version.
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-contact-lists-contact-list-id-contacts
func (c *Client) UserListContactDeleteContacts(ctx context.Context, id, contactListID string) (interface{}, error) {
	r := new(UserContactListContacts)

	return r, c.deleteJSON(ctx, fmt.Sprintf("/users/%s/contact_lists/%s/contacts/", id, contactListID), r)
}

// UserBookmarks gets all the user’s saved events.
// In order to update the saved events list, the user must unsave or save each event.
// A user is authorized to only see his/her saved events.
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-bookmarks
func (c *Client) UserBookmarks(ctx context.Context, id string, req *UserBookmarksRequest) (*UserBookmarksResponse, error) {
	r := new(UserBookmarksResponse)

	return r, c.getJSON(ctx, fmt.Sprintf("/users/%s/bookmarks/", id), req, r)
}

// UserSaveBookmarks adds a new bookmark for the user. Returns {"created": true}.
// A user is only authorized to save his/her own events.
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-post-users-id-bookmarks-save
func (c *Client) UserSaveBookmarks(ctx context.Context, id string, req *UserSaveBookmarkRequest) (interface{}, error) {
	var v interface{}

	return v, c.getJSON(ctx, fmt.Sprintf("/users/%s/bookmarks/save", id), req, v)
}

// UserUnSaveBookmarks removes the specified bookmark from the event for the user. Returns {"deleted": true}.
// A user is only authorized to unsave his/her own events.
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-post-users-id-bookmarks-unsave
func (c *Client) UserUnSaveBookmarks(ctx context.Context, id string, req *UserUnSaveBookmarkRequest) (interface{}, error) {
	var v interface{}

	return v, c.getJSON(ctx, fmt.Sprintf("/users/%s/bookmarks/unsave", id), req, v)
}

// UserAssortments retrieve the assortment for the user
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-assortment
func (c *Client) UserAssortments(ctx context.Context, id string) (*Assortment, error) {
	a := new(Assortment)

	return a, c.getJSON(ctx, fmt.Sprintf("/users/%s/assortment/", id), nil, a)
}

// UserSetAssortments set a user’s assortment and returns the assortment for the specified user.
//
// https://www.eventbrite.com/developer/v3/endpoints/users/#ebapi-get-users-id-assortment
func (c *Client) UserSetAssortments(ctx context.Context, id string, req *UserSetAssortmentRequest) (*Assortment, error) {
	a := new(Assortment)

	return a, c.postJSON(ctx, fmt.Sprintf("/users/%s/assortment/", id), req, a)
}
