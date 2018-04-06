package eventbrite

import (
	"golang.org/x/net/context"
)

// FormatResult is the response structure for available formats
type FormatResult struct {
	Locale  string `json:"locale"`
	Formats []Format
}

// Format is a type of event presentation - e.g. “seminar”, “workshop”, “concert”. Separate from
// category as you can have a “music concert” or a “music seminar”.
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-format
type Format struct {
	// Format ID
	ID string `json:"id"`
	// The format name
	Name string `json:"format"`
	// A shorter name for display in sidebars and other small spaces.
	ShortName string `json:"short_name"`
}

// Formats returns a list of format as formats.
//
// see @https://www.eventbrite.com/developer/v3/endpoints/formats/#ebapi-get-formats
func (c *Client) Formats(ctx context.Context) (*FormatResult, error) {
	res := new(FormatResult)

	return res, c.getJSON(ctx, "/formats", nil, res)
}

// Format gets a format by ID as format.
//
// https://www.eventbrite.com/developer/v3/endpoints/formats/#ebapi-get-formats-id
func (c *Client) Format(ctx context.Context, id string) (*Format, error) {
	res := new(Format)

	return res, c.getJSON(ctx, "/formats/"+id, nil, res)
}
