package eventbrite

import (
	"fmt"

	"golang.org/x/net/context"
)

// CrossEventDiscount is an object representing a discount that a ticket buyer can use. The term “Cross”
// refers to the fact that this kind of discount can be applied to many events at the same time.
//
// There are four types of discounts:
//
//  - Public Discounts, that any user can see on the listing or checkout pages. Only applied to single event discounts.
//  - Coded Discounts, that requires the user to provide a secret code in order to enable them.
//  - Access Codes, that allow the user to access hidden tickets, but cannot provide a discount.
//  - Hold Discounts, that allow the user to unlock or apply discounts to seats defined as hold.
//
// https://www.eventbrite.co.uk/developer/v3/response_formats/event/#ebapi-std:format-cross_event_discount
type CrossEventDiscount struct {
	// The name of the discount (on public discounts) or the code that
	// user should provide in order to activate it (on access codes or coded discounts)
	Code string `json:"code"`
	// One of access, coded, public or hold, indicating the type of discount
	Type string `json:"type"`
	// The code will be usable until this date
	EndDate DateTime `json:"end_date"`
	// The code will be usable until this amount of seconds before the event start
	EndDateRelative int `json:"end_date_relative"`
	// A fixed amount that is applied as a discount. It doesn’t have a currency, it depends on the event’s
	// currency from 0.01 to 99999.99. Only two decimals are allowed. Will be null for an access code
	AmountOff float64 `json:"amount_off"`
	// A percentage discount that will be applied on the ticket display price during the checkout,
	// from 1.00 to 100.00. Only two decimals are allowed. Will be null for an access code
	PercentOff float64 `json:"percent_off"`
	// The number of times this discount can be used, when 0 means “unlimited”
	QuantityAvailable int `json:"quantity_available"`
	// The number of times the discount was used. This is a display only field, it cannot be written
	QuantitySold int `json:"quantity_sold"`
	// The code will be usable since this date
	StartDate DateTime `json:"start_date"`
	// The code will be usable since this amount of seconds before the event start
	StartDateRelative int `json:"start_date_relative"`
	// On single event discounts, the list of IDs of tickets that are part of event_id for wich
	// this discounts applies to. If empty, means “all the tickets of the event”
	TicketClassIds []string `json:"ticket_class_ids"`
	// On single event discounts, the id of the Event this discount applies to. This is empty
	// for cross event discounts
	EventID string `json:"event_id"`
	// The Event for this discount (only for single event discounts)
	Event Event `json:"event"`
	// On cross event discounts, it is the id of the ticket group for which the discount applies to
	TicketGroupID string `json:"ticket_group_id"`
	// The Ticket Group for this discount (only for cross event discounts)
	TicketGroup TicketGroup `json:"ticket_group"`
	// List of IDs of holds this discount can unlock
	HoldIds []string `json:"hold_ids"`
}

// DiscountCreateRequest is the request structure to create a new CrossEventDiscount
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/cross_event_discounts/#ebapi-id1
type DiscountCreateRequest struct {
	// Code used to activate discount
	Code string `json:"discount.code" validate:"required"`
	// One of access, coded, public or hold, indicating the type of discount
	Type string `json:"discount.type"`
	// Fixed reduction amount
	AmountOff float64 `json:"discount.amount_off"`
	// A percentage discount that will be applied on the ticket display price during the checkout,
	// from 1.00 to 100.00. Only two decimals are allowed. Will be null for an access code
	PercentOff float64 `json:"discount.percent_off"`
	// Number of discount uses
	QuantityAvailable int `json:"discount.quantity_available"`
	// Allow use from this date. A datetime represented as a string in Naive Local
	// ISO8601 date and time format, in the timezone of the event
	StartDate DateTime `json:"start_date"`
	// Allow use from this number of seconds before the event starts. Greater than 59 and multiple of 60
	StartDateRelative int `json:"discount.start_date_relative"`
	// Allow use until this date. A datetime represented as a string in Naive Local ISO8601 date
	// and time format, in the timezone of the event
	EndDate DateTime `json:"discount.end_date"`
	// Allow use until this number of seconds before the event starts. Greater than 59 and multiple of 60
	EndDateRelative int `json:"discount.end_date_relative"`
	// IDs of tickets to limit discount to
	TicketClassIds []string `json:"discount.ticket_class_ids"`
	// ID of the event. Only used for single event discounts
	EventID string `json:"discount.event_id"`
	// ID of the ticket group
	TicketGroupID string `json:"discount.ticket_group_id"`
	// IDs of holds this discount can unlock
	HoldIds []string `json:"discount.old_ids"`
}

// DiscountUpdateRequest is the structure to update a CrossEventDiscount
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/cross_event_discounts/#ebapi-id3
type DiscountUpdateRequest struct {
	// Code used to activate discount
	Code string `json:"discount.code" validate:"required"`
	// Fixed reduction amount
	AmountOff float64 `json:"discount.amount_off"`
	// A percentage discount that will be applied on the ticket display price during the checkout,
	// from 1.00 to 100.00. Only two decimals are allowed. Will be null for an access code
	PercentOff float64 `json:"discount.percent_off"`
	// Number of discount uses
	QuantityAvailable int `json:"discount.quantity_available"`
	// Allow use from this date. A datetime represented as a string in Naive Local
	// ISO8601 date and time format, in the timezone of the event
	StartDate DateTime `json:"start_date"`
	// Allow use from this number of seconds before the event starts. Greater than 59 and multiple of 60
	StartDateRelative int `json:"discount.start_date_relative"`
	// Allow use until this date. A datetime represented as a string in Naive Local ISO8601 date
	// and time format, in the timezone of the event
	EndDate DateTime `json:"discount.end_date"`
	// Allow use until this number of seconds before the event starts. Greater than 59 and multiple of 60
	EndDateRelative int `json:"discount.end_date_relative"`
	// IDs of tickets to limit discount to
	TicketClassIds []string `json:"discount.ticket_class_ids"`
	// IDs of holds this discount can unlock
	HoldIds []string `json:"discount.hold_ids"`
}

// DiscountsGet returns the cross_event_discount with the specified :discount_id
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/cross_event_discounts/#ebapi-cross-event-discounts
func (c *Client) DiscountsGet(ctx context.Context, id string) (*CrossEventDiscount, error) {
	d := new(CrossEventDiscount)

	return d, c.getJSON(ctx, fmt.Sprintf("/discounts/%s/", id), nil, d)
}

// DiscountCreate creates a discount. Returns the created cross_event_discount.
//
// The following conditions define the span of the discount’s effect:
//
//  - If event_id is provided and ticket_class_ids is not provided, a single-event discount for all the tickets in the event is created.
//  - If both event_id and ticket_class_ids are provided, a single-event discount for the specified event tickets is created.
//  - If ticket_group_id is provided, a cross-event discount for the specified ticket group is created.
//  - If neither event_id nor ticket_group_id are provided, a discount that applies to all the events and all tickets of the user is created. This means that the discount will apply to future events also.
//
// Notes:
//
// Public and coded discounts can have either an amount off or a percentage off, but not both. Access codes cannot have an amount or percentage off.
// Public discounts should not contain apostrophes or non-alphanumeric characters (except “-”, “_”, ” ”, “(”, ”)”, “/”, and “”).
// Coded discounts and access codes should not contain spaces, apostrophes or non-alphanumeric characters (except “-”, “_”, “(”, ”)”, “/”, and “”).
//
//  - If the start_date and start_date_relative are null or empty, that means that the discount is usable effective immediately.
//  - If the end_date and end_date_relative are null or empty, that means that the discount is usable until the event finishes.
//  - If start_date_relative is provided, the discount will be usable after the given number of seconds prior to the event start.
//  - If end_date_relative is provided, the discount will be usable until the given number of seconds prior to the event start.
//
// Discount for series events should be associated with the parent event, not its children
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/cross_event_discounts/#ebapi-post-discounts
func (c *Client) DiscountCreate(ctx context.Context, req *DiscountCreateRequest) (*CrossEventDiscount, error) {
	d := new(CrossEventDiscount)

	return d, c.postJSON(ctx, "/discounts/", req, d)
}

// DiscountUpdate updates the discount with the specified :discount_id. Returns the updated cross_event_discount.
// The fields sent are the ones that are going to be updated, the fields that are not sent will be
// unchanged. The same conditions and notes for the discount creation apply
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/cross_event_discounts/#ebapi-post-discounts-discount-id
func (c *Client) DiscountUpdate(ctx context.Context, id string, req *DiscountUpdateRequest) (*CrossEventDiscount, error) {
	d := new(CrossEventDiscount)

	return d, c.postJSON(ctx, fmt.Sprintf("/discounts/%s/", id), req, d)
}

// DiscountDelete deletes the cross_event_discount with the specified :discount_id. Only unused discounts can be deleted.
// Warning: The discount cannot be restored after deletion.
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/cross_event_discounts/#ebapi-delete-discounts-discount-id
func (c *Client) DiscountDelete(ctx context.Context, id string) (interface{}, error) {
	var v interface{}

	return v, c.deleteJSON(ctx, fmt.Sprintf("/discounts/%s/", id), &v)
}
