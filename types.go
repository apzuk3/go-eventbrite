package eventbrite

import (
	"time"
	"bytes"
	"fmt"
)

// When an error occurs during an API request, you’ll get a response with an error HTTP status
// (in the 400 or 500 range), as well as a JSON response containing more information about the error.
//
// https://www.eventbrite.co.uk/developer/v3/api_overview/errors/#ebapi-errors
type Error struct {

	// The error key contains a constant string value for error - in this case, VENUE_AND_ONLINE - and
	// is what you should key your error handling off of, as this string won’t change depending on locale
	// or as we change the API over time
	Err string `json:"error"`

	// The error_description key is for developer information only and will usually contain a more informative
	// explanation for the error, should you be confused. You should not display this string to your users;
	// it’s often very technical and may not be localized to their language
	Description string `json:"error_description"`

	// The status_code value just mirrors the HTTP status code you got as part of the request. It’s there as
	// a convenience if your HTTP library makes it very hard to get status codes, or has one error handler
	// for all error codes
	Status int `json:"status_code"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Eventbrite API: [Status code - %d] %s", e.Status, e.Description)
}

// The ISO 3166 alpha-2 code of a country.
type CountryCode string

// An ISO 4217 3-character code of a currency
type CurrencyCode string

type Currency struct {
	Currency CurrencyCode `json:"currency"`
	Value float32 `json:"value"`
	Display string `json:"display"`
}

type Date struct {
	Time time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	data = bytes.Replace(data, []byte("\""), []byte(""), -1)
	t, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		fmt.Println(err)
	}

	d.Time = t
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Time.Format("2006-01-02") + "\""), nil
}

type DateTime struct {
	Time time.Time
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	data = bytes.Replace(data, []byte("\""), []byte(""), -1)
	t, err := time.Parse("2006-01-02T15:04:05Z", string(data))
	if err != nil {
		fmt.Println(err)
	}

	d.Time = t
	return err
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Time.Format("2006-01-02T15:04:05Z") + "\""), nil
}

// Timezone is an object with details about a timezone
type Timezone struct {

	// Timezone id
	Id string `json:"id"`

	// The timezone identifier as defined by the IANA Time Zone Database
	Timezone string `json:"timezone"`

	// The localized name for the timezone
	Label string `json:"label"`

	ResourceUri string `json:"resource_uri"`
}

// All API endpoints which return multiple objects will return paginated responses; as well as the
// list of objects (which will usually be under a key like events or attendees, depending on the endpoint)
// there will also be a pagination key:
//
// see @https://www.eventbrite.com/developer/v3/api_overview/pagination/#ebapi-paginated-responses
type Pagination struct {
	ObjectCount int `json:"object_count"`
	PageNumber int `json:"page_number"`
	PageSize int `json:"page_size"`
	PageCount int `json:"page_count"`
	HasMoreItems bool `json:"has_more_items"`
}


// Returned for fields which represent HTML, like event names and descriptions.
// The html key represents the original HTML (which _should_ be sanitized and free from injected script tags etc.,
// but as always, be careful what you put in your DOM), while the text key is a stripped version useful for places
// where you can’t or don’t need to display the full HTML version.
//
// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-multipart-text
type MultipartText struct {
	Text string `json:"text"`
	Html string `json:"html"`
}

// A combination of a timezone from the Olson specification as a string, and two datetime values, one for
// the UTC time represented and one for the local time in the named timezone.
//
// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-datetime-with-timezone
type DatetimeTz struct {
	Timezone string `json:"timezone"`
	Utc string `json:"utc"`
	Local string `json:"local"`
}

// Country is an object with details about a country
//
// https://www.eventbrite.com/developer/v3/response_formats/system/#ebapi-countries
type Country struct {
	// The country identifier as defined by the ISO 3166 standard
	Code CountryCode `json:"code"`

	// The readable name of the country
	Label string `json:"label"`
}

// Region is an object with details about a region
//
// https://www.eventbrite.com/developer/v3/response_formats/system/#ebapi-region
type Region struct {

	// The associated country code to this region
	CountryCode string `json:"country_code"`

	// The region identifier as defined by the ISO 3166 standard
	Code string `json:"code"`

	// The readable name of the region
	Label string `json:"label"`
}

// Image is an object with details about a given image.
//
// https://www.eventbrite.com/developer/v3/response_formats/image/#ebapi-image
type Image struct {

	// The image’s ID
	Id string `json:"id"`

	// The URL of the image
	Url string `json:"url"`
}

// A location where an event happens.
//
// https://www.eventbrite.com/developer/v3/response_formats/venue/#ebapi-venue
type Venue struct {

	// The value name
	Name string `json:"name"`

	// The address of the venue
	Address Address `json:"name"`
}

// Though address formatting varies considerably between different countries and regions, Eventbrite
// still has a common address return format to keep things consistent.
//
// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-address
type Address struct {

	// The street/location address (part 1)
	Address1 string `json:"address_1"`

	// The street/location address (part 2)
	Address2 string `json:"address_2"`

	// The city
	City string `json:"city"`

	// The ISO 3166-2 2- or 3-character region code for the state, province, region, or district
	Region string `json:"region"`

	// The postal code
	PostalCode string `json:"postal_code"`

	// The ISO 3166-1 2-character international code for the country
	Country string `json:"country"`

	// The latitude portion of the address coordinates
	Latitude string `json:"latitude"`

	// The longitude portion of the address coordinates
	Longitude string `json:"longitude"`

	// The format of the address display localized to the address country
	LocalizedAddressDisplay string `json:"localized_address_display"`

	// The format of the address’s area display localized to the address country
	LocalizedAreaDisplay string `json:"localized_area_display"`

	// 	The multi-line format order of the address display localized to the address country, where each line is an item in the list
	LocalizedMultiLineAddressDisplay []interface{} `json:"localized_multi_line_address_display"`
}

// A grouping entity that Eventbrite uses to display as the owner of events. Contains name
// and contact information.
//
// https://www.eventbrite.com/developer/v3/response_formats/organizer/#ebapi-std:format-organizer
type Organizer struct {

	// The organizer name
	Name string `json:"name"`

    // The description of the organizer (may be very long and contain significant formatting)
	Description MultipartText `json:"description"`

	// The URL to the organizer’s page on Eventbrite
	Url string `json:"url"`
}


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
	Venue Venue `json:"venue"`
	VenueId string `json:"venue_id"`

	// The organizer of the event
	Organizer Organizer `json:"organizer"`
	OrganizerId string `json:"organizer_id"`

	// The event’s format (type of event: conference, seminar, concert, etc.) (optional)
	Format Format `json:"format"`
	FormatId string `json:"format_id"`

	// The event’s category (technology, music, science, etc.) (optional)
	Category Category `json:"category"`
	CategoryId string `json:"category_id"`

	// The event’s subcategory (optional)
	SubCategory SubCategory `json:"subcategory"`
	SubCategoryId string `json:"subcategory_id"`

	Logo Image `json:"logo"`

	// The event’s refund policy (optional)
	RefundPolicy interface{} `json:"refund_policy"`

	// The bookmark information on the event. Currently returns a dictionary with the number of users who
	// have bookmarked the event as ‘count’ (i.e. {'count': 3})
	BookmarkInfo interface{} `json:"bookmark_info"`
}

// An overarching category that an event falls into (vertical). Examples are “Music”, and “Endurance”.
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-category
type Category struct {

	// Category ID
	ID string `json:"id"`

	// he category name
	Name string `json:"name"`

	// The category name localized to the current locale (if available)
	NameLocalized string `json:"name_localized"`

	// A shorter name for display in sidebars and other small spaces.
	ShortName string `json:"sort_name"`

	// List of subcategories, only shown on some endpoints.
	ShortNameLocalized string `json:"short_name_localized"`

	SubCategories []SubCategory

	ResourceUri string `json:"resource_uri"`
}

// A more specific category that an event falls into, sitting underneath a category.
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-subcategory
type SubCategory struct {

	// Subcategory ID
	ID string `json:"id"`

	// The category name
	Name string `json:"name"`

	// The category this belongs to
	ParentCategory Category `json:"parent_category"`
}

// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-display-settings
type EventSettings struct {

	// Whether to display the start date on the event listing
	ShowStartDate bool `json:"display_settings.show_start_date"`

	// 	Whether to display the end date on the event listing
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

// This is an object that details what fees are applied for a specific set of conditions.
//
// https://www.eventbrite.com/developer/v3/response_formats/pricing/#ebapi-std:format-fee_rate
type FeeRate struct {

	// The (ISO 3166 alpha-2 code of the) country
	Country CountryCode `json:"country"`

	// The (ISO 4217 3-character code of the) currency
	Currency CurrencyCode `json:"currency"`

	// The assortment package name to get the price for, one of (‘any’, ‘package1’, ‘package2’).
	// ‘any’ means that applies to all the prossible variants.
	Plan string `json:"place"`

	// The payment type to get the price for, one of (‘any’, ‘eventbrite’, ‘authnet’, ‘moneris’,
	// ‘paypal’, ‘google’, ‘manual’, ‘free’, ‘offline’, ‘cash’, ‘check’, ‘invoice’). ‘any’
	// means that applies to all the prossible variants
	PaymentType string `json:"payment_type"`

	// The sales channel. One of (‘any’, ‘atd’, ‘online’). ‘any’ means that applies to all the prossible variants
	Channel string `json:"channel"`

	// The item type for which get the price fee rates. One of (‘any’, ‘ticket’, ‘product’). ‘any’ means that
	// applies to all the prossible variants
	ItemType string `json:"item_type"`

	// FeeRate rule percent. Minimum value is ‘0’, maximum value is ‘100’. Supports two decimals
	Percent float32 `json:"percent"`

	// Name of the fee (service_fee or payment_fee).
	Name string `json:"fee_name"`

	// FeeRate rule fixed value
	Fixed Currency `json:"fixed"`

	// FeeRate rule maximum amount (Cap). Null means unlimited
	Maximum Currency `json:"maximum"`

	// FeeRate rule minimum amount. Null means that there isn’t any minimum
	Minimum Currency `json:"minimum"`
}

// Contains a refund request of the order
//
// https://www.eventbrite.com/developer/v3/response_formats/order/#ebapi-std:format-refund-request
type RefundRequest struct {

	// The email used to create the refund request
	FromEmail string `json:"from_email"`

	// The name used to create the refund request
	FromName string `json:"from_name"`

	// The actual status of the refund request
	Status string `json:"status"`

	// The message associated with the refund request
	Message string `json:"message"`

	// The code of the refund request’s reason
	Reason string `json:"reason"`

	// The last message associated with the last status of the refund request
	LastMessage string `json:"last_message"`

	// The last code of the refund request’s reason
	LastReason string `json:"last_reason"`

	// The items of the refund request
	Items []RefundItem `json:"items"`
}

// Contains a refund item
//
// https://www.eventbrite.com/developer/v3/response_formats/order/#ebapi-refund-item
type RefundItem struct {

	// The event of this item
	EventId string `json:"event_id"`

	// The order of this item. Consider that this field can be null
	OrderId string `json:"order_id"`

	// the item type order for full refund, attendee for partial refund an
	// attendee or merchandise for partial refund a merchandise
	ItemType string `json:"item_type"`

	// The quantity requested for this item. If the item_type is order, quantity_requested is always 1.
	// if the item_type is attendee or merchandise, then the quantity_requested shows how
	// many items were requested
	QuantityRequested int `json:"quantity_requested"`

	// The total amount requested for this item.
	AmountRequested Currency `json:"amount_requested"`
}

// The Ticket Group object allows the users to group an arbitrary number of ticket_class
//
// https://www.eventbrite.com/developer/v3/response_formats/ticket_group/#ebapi-std:format-ticket_group
type TicketGroup struct {

	Id string `json:"id"`

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

// This is an object representing one of the possible ticket classes (types of ticket) for an event
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-ticket-class
type TicketClass struct {

	Id string `json:"id"`

	// The ticket class’ name
	Name string `json:"name"`

	// The ticket’s description. (optional)
	Description string `json:"description"`

	// The display cost of the ticket (paid only)
	Cost Currency `json:"cost"`

	// The display fee of the ticket (paid only)
	Fee Currency `json:"fee"`

	// If the ticket is a donation
	Donation bool `json:"donation"`

	// If the ticket is a free ticket
	Free bool `json:"free"`

	// Minimum number that can be bought per order
	MinimumQuantity int `json:"minimum_quantity"`

	// Maximum number that can be bought per order
	MaximumQuantity int `json:"maximum_quantity"`

	// The event the ticket class is for
	EventId string `json:"event_id"`

	// The event the ticket class is for
	Event Event `json:"event"`

	// PRIVATE FIELDS
	// Only shown to people with event owner permission

	// How many of these tickets are available to be sold overall
	QuantityTotal int `json:"quantity_total"`

	// How many of these tickets have already been sold and confirmed (does not include tickets being checked out right now)
	QuantitySold int `json:"quantity_sold"`

	// If the ticket is hidden from the public
	Hidden bool `json:"hidden"`

	// When sales for this ticket start
	SalesStart string `json:"sales_start"`

	// When sales for this ticket end
	SalesEnd string `json:"sales_end"`

	// The ID of another ticket class that, when it sells out, will trigger sales of this class to start
	SalesStartAfter string `json:"sales_start_after"`

	// If the fee should be included in the displayed cost (cannot be set along with split_fee)
	IncludeFee bool `json:"include_fee"`

	// If the payment fee should be included in the displayed cost and the eventbrite fee is shown separately
	SplitFee bool `json:"split_fee"`

	// If the description should be hidden on the event page (will remove description from public responses too)
	HideDescription bool `json:"hide_description"`

	// If the ticket should be hidden when not on sale
	AutoHide bool `json:"auto_hide"`

	// Override the time at which auto hide disables itself to show the ticket (otherwise it’s sales_start)
	AutoHideBefore string `json:"auto_hide_before"`

	// Override the time at which auto hide enables itself to re-hide the ticket (otherwise it’s sales_end)
	AutoHideAfter  string `json:"auto_hide_after"`
}

// An entity that Eventbrite uses to allow event organizer to utilize tracking pixels on their events
//
// https://www.eventbrite.com/developer/v3/response_formats/tracking_beacon/#ebapi-tracking-beacon
type TrackingBeacon struct {

	// The tracking beacon id
	Id string

	// The tracking beacon third party type. Allowed types are: Facebook Pixel,
	// Twitter Ads, AdWords, Google Analytics, Simple Image Pixel, Adroll iPixel
	TrackingType string

	// The id of the event where the tracking beacon will load your tracking pixel
	EventId string

	// The id of the user where the tracking beacon will load this tracking pixel on all of their events
	UserId string

	// The third party id that they have given you to fire on your event page
	PixelId string

	// The tracking pixel meta information that determines where your pixel will fire
	Triggers interface{}
}

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

// Contains a list of email objects giving information on the user’s email addresses
//
// https://www.eventbrite.com/developer/v3/response_formats/user/#ebapi-user-emails
type Email struct {
	Email string `json:"email"`
	Verified bool `json:"verified"`
	Primary bool `json:"primary"`
}

// An object representing a single webhook associated with the account
type Webhook struct {

	// The uri of the individual webhook
	ResourceUri string `json:"resource_uri"`

	// The url that the webhook will send data to when it is triggered
	EndpointUrl string `json:"endpoint_url"`

	// One or any combination of actions that will cause this webhook to fire
	Actions string `json:"actions"`
}

// Attendee is an object representing the details of one or more people coming to the event
// Attendee objects are considered private and are only available to the event owner
type Attendee struct {

	// When the attendee was created (order placed)
	Created DateTime `json:"created"`

	// When the attendee was last changed
	Changed DateTime `json:"changed"`

	// The name of the ticket_class at the time of registration
	TicketClassName string `json:"ticket_class_name"`

	// The attendee’s basic profile information
	Profile AttendeeProfile `json:"profile"`

	// The attendee’s basic profile information
	Addresses AttendeeAddresses `json:"addresses"`

	// The attendee’s answers to any custom questions (optional)
	Answers AttendeeAnswers `json:"answers"`

	// The attendee’s entry barcode information
	Barcodes AttendeeBarcodes `json:"barcodes"`

	// The attendee’s team information (optional)
	Team AttendeeTeam `json:"team"`

	// The attendee’s affiliate code (optional)
	//
	// Not documented
	Affiliate interface{} `json:"affiliate"`

	// If the attendee is checked in
	CheckedIn bool `json:"checked_in"`

	// If the attendee is cancelled
	Cancelled bool `json:"cancelled"`

	// If the attendee is refunded
	Refunded bool `json:"refunded"`

	// The status of the attendee (scheduled to be deprecated)
	Status string `json:"status"`

	// The event id that this attendee is attending
	EventID string `json:"event_id"`

	// The event this attendee is attending
	Event Event `json:"event"`

	// The order id this attendee is part of
	OrderID string `json:"order_id"`

	// The order this attendee is part of
	Order Order `json:"order"`

	// The guestlist id for this attendee. If this is null it means that this is not a guest
	GuestListID string `json:"guestlist_id"`

	// The guest of for the guest. If this is null it means that this is not a guest
	InvitedBy string `json:"invited_by"`

	// The promotional code applied to this attendee
	//
	// Not documented
	PromotionalCode interface{} `json:"promotional_code"`

	// The bib number assigned to this attendee if one exists for a race or endurance event
	//
	// Not documented
	AssignedNumber interface{} `json:"assigned_number"`
}

// Contains the attendee’s personal information
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-std:format-attendee-profile
type AttendeeProfile struct {

	// The attendee’s name. Use this in preference to first_name/last_name/etc. if possible for
	// forward compatibility with non-Western names
	Name string `json:"name"`

	// The attendee’s email address
	Email string `json:"email"`

	// The attendee’s first name
	FirstName string `json:"first_name"`

	// The attendee’s last name
	LastName string `json:"last_name"`

	// The title or honoraria used in front of the name (Mr., Mrs., etc.) (optional)
	Prefix string `json:"prefix"`

	// The suffix at the end of the name (e.g. Jr, Sr) (optional)
	Suffix string `json:"suffix"`

	// The attendee’s age (optional)
	Age int `json:"age"`

	// The attendee’s job title (optional)
	JobTitle string `json:"job_title"`

	// The attendee’s company name (optional)
	Company string `json:"company"`

	// The attendee’s website address (optional)
	Website string `json:"website"`

	// The attendee’s blog address (optional)
	Blog string `json:"blog"`

	// The attendee’s gender (currently one of “male” or “female”) (optional)
	Gender string `json:"gender"`

	// The attendee’s birth date (optional)
	BirthDate Date `json:"birth_date"`

	// The attendee’s cell/mobile phone number, as formatted by them (optional)
	CellPhone string `json:"cell_phone"`
}

// Contains the attendee’s various different addresses. All are optional
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-addresses
type AttendeeAddresses struct {

	// The attendee’s home address
	Home Address `json:"home"`

	// The attendee’s ship address
	Ship Address `json:"ship"`

	// The attendee’s workl address
	Work Address `json:"work"`
}

// A list of objects with answers to custom questions
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-answers
type AttendeeAnswers struct {

	// The ID of the custom question
	QuestionID string `json:"question_id"`

	// The text of the custom question
	Question string `json:"question"`

	// One of multiple_choice, or text
	Type string `json:"type"`

	// The attendee’s answer
	Answer string `json:"answer"`
}

// A list of objects representing the barcodes for this order (usually only one per attendee)
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-barcodes
type AttendeeBarcodes struct {

	//  The barcode contents. Note that if the event organizer has turned off printable
	// tickets, this field will be null in order to prevent exposing the barcode value
	Barcode string `json:"barcode"`

	// One of unused, used, or refunded
	Status string `json:"status"`

	// When the attendee barcode was created
	Created DateTime `json:"created"`

	// When the attendee barcode was changed
	Changed DateTime `json:"changed"`
}

// Represents team information for the attendee if the event has teams configured
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-team
type AttendeeTeam struct {

	// The team’s ID
	ID string `json:"id"`

	// The team’s name
	Name string `json:"name"`

	// When the attendee joined the team
	DateJoined DateTime `json:"date_joined"`

	// The event the team is part of
	EventID string `json:"event_id"`
}