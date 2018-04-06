package eventbrite

import "golang.org/x/net/context"

type Timezones struct {
	Locale     string     `json:"locale"`
	Timezones  []Timezone `json:"timezones"`
	Pagination Pagination `json:"pagination"`
}

type Regions struct {
	Locale     string     `json:"locale"`
	Regions    []Region   `json:"regions"`
	Pagination Pagination `json:"pagination"`
}

type Countries struct {
	Locale     string     `json:"locale"`
	Countries  []Country  `json:"countries"`
	Pagination Pagination `json:"pagination"`
}

// Timezones returns a paginated response with a key of timezones, containing a list of timezones
//
// https://www.eventbrite.com/developer/v3/endpoints/system/#ebapi-get-system-timezones
func (c *Client) Timezones(ctx context.Context) (*Timezones, error) {
	res := new(Timezones)

	return res, c.getJSON(ctx, "/system/timezones/", nil, res)
}

// Timezones returns a single page response with a key of regions, containing a list of regions
//
// https://www.eventbrite.com/developer/v3/endpoints/system/#ebapi-get-system-regions
func (c *Client) Regions(ctx context.Context) (*Regions, error) {
	res := new(Regions)

	return res, c.getJSON(ctx, "/system/regions/", nil, res)
}

// Timezones returns a single page response with a key of countries, containing a list of countries
//
// https://www.eventbrite.com/developer/v3/endpoints/system/#ebapi-get-system-countries
func (c *Client) Countries(ctx context.Context) (*Countries, error) {
	res := new(Countries)

	return res, c.getJSON(ctx, "/system/regions/", nil, res)
}
