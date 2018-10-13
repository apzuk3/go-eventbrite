package eventbrite

import (
	"fmt"

	"golang.org/x/net/context"
)

// Checkout is an object that represents the settings for how an organizer
// wants ticket buyers pay for their purchases.
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-checkout-settings-countries-currencies
type Checkout struct {

	// a list of supported ISO 3166-1 2-letter countries
	Countries []string `json:"countries"`

	// a list of supported ISO 4217 3-letter currencies
	Currencies []string `json:"currencies"`

	// a map of ISO 3166-1 alpha-2 country codes to their default ISO 4217 3-letter currency code
	DefaultCurrenciesByCountry map[string]string `json:"currencies"`
}

// CheckoutMethodsResponse is the response structure for the
// available checkout methods to do payments given a country and a currency
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-checkout-settings-methods
type CheckoutMethodsResponse struct {

	// a list with supported checkout methods given a country and currency combination.
	// Set of possible values: [authnet, eventbrite, offline, paypal]
	Methods []string `json:"methods"`
}

// CheckoutSettingsForAccount is the response structure of Checkout settings for the current
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-checkout-settings
type CheckoutSettingsForAccount struct {
	CheckoutSettings []Checkout `json:"checkout_settings"`
}

// CheckoutMethodsRequest is the request structure for the available
// checkout methods to do payments given a country and a currency
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-id1
type CheckoutMethodsRequest struct {

	// Expected methods for Country
	Country string `json:"country" validate:"required"`

	// Expected methods for Currency
	Currency string `json:"currency" validate:"required"`
}

// CheckoutForAccountRequest is the request to search Checkout settings for the current user
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-id3
type CheckoutForAccountRequest struct {
	// An optional country code by which to filter checkout settings
	Country string `json:"country"`

	// An optional currency code by which to filter checkout settings
	Currency string `json:"currency"`

	// One or more optional (comma-separated) checkout methods by which to filter checkout settings
	CheckoutMethods string `json:"checkout_methods"`

	SearchMostRecentEvent bool `json:"search_most_recent_event"`
}

// CheckoutCreateRequest is the request structure for creating a new Checkout settings
// object belonging to the current user
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-id5
type CheckoutCreateRequest struct {

	// The country code for the checkout settings object
	CountryCode string `json:"checkout_settings.country_code" validate:"required"`

	// The currency code for the checkout settings object
	CurrencyCode string `json:"checkout_settings.currency_code" validate:"required"`

	// The checkout method for the checkout settings object
	Method string `json:"checkout_settings.checkout_method" validate:"required"`

	// The vault ID for the user instrument if the checkout method requires one
	UserInstrumentVaultID string `json:"checkout_settings.user_instrument_vault_id"`

	// A list of additional settings for the offline checkout method, with each offline setting being in the
	// format {"payment_method": "CASH"|"CHECK"|"INVOICE", "instructions": "Optional instructions"}. Required
	// if the checkout_method is “offline.”
	// Example:
	//  [
	//    {
	//       "payment_method": "CASH"
	//    },
	//    {
	//       "payment_method": "CHECK",
	//       "instructions": "Make checks payable to ABC corporation"
	//    },
	//    ...
	//  ]
	//
	// also https://www.eventbrite.co.uk/developer/v3/response_formats/basic/#ebapi-std:format-objectlist
	OfflineSettings interface{} `json:"checkout_settings.offline_settings"`

	// For the “paypal” checkout method, you can optionally specify a PayPal account email address instead
	// of a user instrument vault ID, and a matching user instrument will be found or a new user instrument
	// created with that email address and used to create the checkout settings.
	PaypalEmail string `json:"paypal_email"`
}

// CheckoutAssociateToEventRequest is the request structure to associate
// a single or set of Checkout seeting with a given event by its event_id
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-id12
type CheckoutAssociateToEventRequest struct {

	// A list of IDs for checkout settings that should be linked to the event. In the format: 1234,5678,9012
	CheckoutSettingsIds []string `json:"checkout_settings_ids"`
}

// CheckoutAssociatePayoutToEvent is the request structure to associate a payout
// user instrument ID with a given event, or clear the association by passing a
// null value for the user instrument ID
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-id17
type CheckoutAssociatePayoutToEvent struct {

	// The vault ID for the user instrument to which payouts are sent
	UserInstrumentVaultID string `json:"payout_settings.user_instrument_vault_id"`
}

// CheckoutGetList gets the countries and currencies which are supported by Eventbrite for ticket payment
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-checkout-settings-countries-currencies
func (c *Client) CheckoutGetList(ctx context.Context) (*Checkout, error) {
	s := new(Checkout)

	return s, c.getJSON(ctx, "/checkout_settings/countries_currencies/", nil, s)
}

// CheckoutMethods gets the available checkout methods to do payments given a country and a currency
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-checkout-settings-methods
func (c *Client) CheckoutMethods(ctx context.Context, req CheckoutMethodsRequest) (*CheckoutMethodsResponse, error) {
	s := new(CheckoutMethodsResponse)

	return s, c.getJSON(ctx, "/checkout_settings/methods/", nil, s)
}

// CheckoutForAccount searches and returns a list of checkout_settings for the current
// user as the key checkout_settings
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-checkout-settings
func (c *Client) CheckoutForAccount(ctx context.Context, req *CheckoutForAccountRequest) (*CheckoutSettingsForAccount, error) {
	s := new(CheckoutSettingsForAccount)

	return s, c.getJSON(ctx, "/checkout_settings/", nil, s)
}

// CheckoutCreate creates a new checkout_settings object belonging to the current user. Two
// common settings are Eventbrite. Payment Processing ( checkout_method = “eventbrite” )
// and PayPal ( checkout_method = “paypal” ). In addition to the checkout_method you must
// provide the country and currency proceeds from the event should be paid to
//
// For all checkout methods except “eventbrite” and “offline” you must provide a valid user_instrument_vault_id
//
// Returns a list of checkout_settings
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-post-checkout-settings
func (c *Client) CheckoutCreate(ctx context.Context, req *CheckoutCreateRequest) (*Checkout, error) {
	s := new(Checkout)

	return s, c.postJSON(ctx, "/checkout_settings/", req, s)
}

// CheckoutGet gets a specific checkout_settings object by ID
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-checkout-settings-checkout-settings-id
func (c *Client) CheckoutGet(ctx context.Context, id string) (*Checkout, error) {
	s := new(Checkout)

	return s, c.getJSON(ctx, fmt.Sprintf("/checkout_settings/%s/", id), nil, s)
}

// CheckoutByEvent gets and returns a list of checkout_settings associated with a given event by its event_id
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-get-events-event-id-checkout-settings
func (c *Client) CheckoutByEvent(ctx context.Context, eventId string) ([]*Checkout, error) {
	var s []*Checkout

	return s, c.getJSON(ctx, fmt.Sprintf("/events/%s/checkout_settings/", eventId), nil, s)
}

// CheckoutAssociate associates a single or set of checkout_settings with a given event by its event_id. This does not add
// more checkout settings to the event, but instead replaces all checkout settings for the event with
// the one(s) submitted. The JSON post body is a string list of the checkout_settings IDs you want to associate
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-post-events-event-id-checkout-settings
func (c *Client) CheckoutAssociate(ctx context.Context, eventID string, req *CheckoutAssociateToEventRequest) (interface{}, error) {
	var v interface{}

	return v, c.postJSON(ctx, fmt.Sprintf("/events/%s/checkout_settings/", eventID), req, v)
}

// Associates a payout user instrument ID with a given event, or clear the association by
// passing a null value for the user instrument ID
//
// https://www.eventbrite.co.uk/developer/v3/endpoints/checkout_settings/#ebapi-post-events-event-id-payout-settings
func (c *Client) CheckoutAssociatePayoutSettings(
	ctx context.Context,
	eventID string,
	req *CheckoutAssociatePayoutToEvent) (interface{}, error) {
	var v interface{}

	return v, c.postJSON(ctx, fmt.Sprintf("/events/%s/checkout_settings/", eventID), req, v)

}
