package eventbrite

import "golang.org/x/net/context"

// FeeRate is an object that details what fees are applied for a specific set of conditions.
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

// Returns a list of fee_rate objects for the different currencies, countries, assortments
// and sales channels we sell through today and in the future.
//
// https://www.eventbrite.com/developer/v3/endpoints/pricing/#ebapi-get-pricing-fee-rates
type FeeRequest struct {
	// The (ISO 3166 alpha-2 code of the) country where you want to know the fee rates
	Country CountryCode `json:"country" validate:"required"`
	//     The (ISO 4217 3-character code of the) currency where you want to know the fee rates
	Currency CurrencyCode `json:"currency" validate:"required"`
	// The assortment package name to get the price for. One of [‘any’, ‘package1’, ‘package2’].
	// If it’s not provided, or the value is ‘any’, all the existing variants will be returned.’
	Plan string `json:"plan"`
	// The payment type to get the price for. One of [‘any’, ‘eventbrite’, ‘authnet’, ‘moneris’,
	// ‘paypal’, ‘google’, ‘manual’, ‘free’, ‘offline’, ‘cash’, ‘check’, ‘invoice’]. If it’s not provided,
	// or the value is ‘any’, all the existing variants will be returned.
	PaymentType string `json:"payment_type"`
	// The sales channel. One of [‘any’, ‘atd’, ‘web’]. If it’s not provided, or the value is ‘any’,
	// all the existing variants will be returned.
	Channel string `json:"channel"`
	// The item type for which get the price fee rates. One of [‘any’, ‘ticket’, ‘product’]. If it’s not provided,
	// or the value is ‘any’, all the existing variants will be returned.
	ItemType string `json:"item_type"`
}

// FeeResponse is the response structure for fee rate request
//
// https://www.eventbrite.com/developer/v3/endpoints/pricing/#ebapi-get-pricing-fee-rates
type FeeResponse struct {
	FeeRates []FeeRate `json:"fee_rates"`
}

// FeeRate returns a list of fee_rate objects for the different currencies, countries, assortments
// and sales channels we sell through today and in the future.
//
// https://www.eventbrite.com/developer/v3/endpoints/pricing/#ebapi-get-pricing-fee-rates
func (c *Client) FeeRate(ctx context.Context, req *FeeRequest) (*FeeResponse, error) {
	res := new(FeeResponse)

	return res, c.getJSON(ctx, "/pricing/fee_rates", req, res)
}
