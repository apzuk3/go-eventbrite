package eventbrite

import "golang.org/x/net/context"

// https://www.eventbrite.com/developer/v3/endpoints/reports/#ebapi-parameters
type ReportRequest struct {
	// List of public event IDs to report on
	EventIds []interface{} `json:"event_ids"`
	// Event status to filter down results by (Valid choices are: all, live, or ended)
	EventStatus string `json:"event_status"`
	// Optional start date to query
	StartDate string `json:"start_date"`
	// Optional end date to query
	EndDate string `json:"end_date"`
	// Time period to provide aggregation for in units of the selected date_facet.
	// For example, if date_facet=hour, then period=3 returns 3 hours worth of data
	// from the current time in the event timezone. Day is the default choice if no date_facet
	Period int `json:"period"`
	// Optional filters for sales/attendees data formatted as: {“ticket_ids”: [1234, 5678],
	// “currencies”: [“USD”],...}NOTE: currently only filter_by ticket_ids and one currency are supported.
	//
	// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-dictionary
	FilterBy interface{} `json:"filter_by"`
	// Optional field to group data on (Valid choices are: payment_method, payment_method_application,
	// ticket, ticket_application, currency, event_currency, reserved_section, event, event_ticket,
	// event_application, country, city, state, source, zone, location, access_level, device_name,
	// sales_channel_lvl_1, sales_channel_lvl_2, or sales_channel_lvl_3)
	GroupBy string `json:"group_by"`
	// Optional date aggregation level to return data for. Day is the default choice. Monthly aggregation
	// is represented by the first of the month. Weekly aggregation is represented by the ending Sunday of
	// the week, where a week is defined as Monday-Sunday. (Valid choices are: fifteen, hour, day, event_day,
	// week, month, year, or none)
	DateFacet string `json:"date_facet"`
	// Optional timezone. If unspecified picks the TZ of the first event
	Timezone string `json:"timezone"`
}

// https://www.eventbrite.com/developer/v3/endpoints/reports/#ebapi-id1
type ReportAttendees struct {
	// List of public event IDs to report on
	EventIds []interface{} `json:"event_ids"`
	// Event status to filter down results by (Valid choices are: all, live, or ended)
	EventStatus string `json:"event_status"`
	// Optional start date to query
	StartDate string `json:"start_date"`
	// Optional end date to query
	EndDate string `json:"end_date"`
	// Time period to provide aggregation for in units of the selected date_facet.
	// For example, if date_facet=hour, then period=3 returns 3 hours worth of data
	// from the current time in the event timezone. Day is the default choice if no date_facet
	Period int `json:"period"`
	// Optional filters for sales/attendees data formatted as: {“ticket_ids”: [1234, 5678],
	// “currencies”: [“USD”],...}NOTE: currently only filter_by ticket_ids and one currency are supported.
	//
	// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-dictionary
	FilterBy interface{} `json:"filter_by"`
	// Optional field to group data on (Valid choices are: payment_method, payment_method_application,
	// ticket, ticket_application, currency, event_currency, reserved_section, event, event_ticket,
	// event_application, country, city, state, source, zone, location, access_level, device_name,
	// sales_channel_lvl_1, sales_channel_lvl_2, or sales_channel_lvl_3)
	GroupBy string `json:"group_by"`
	// Optional date aggregation level to return data for. Day is the default choice. Monthly aggregation
	// is represented by the first of the month. Weekly aggregation is represented by the ending Sunday of
	// the week, where a week is defined as Monday-Sunday. (Valid choices are: fifteen, hour, day, event_day,
	// week, month, year, or none)
	DateFacet string `json:"date_facet"`
	// Optional timezone. If unspecified picks the TZ of the first event
	Timezone string `json:"timezone"`
}

// ReportSales returns a response of the aggregate sales data
//
// https://www.eventbrite.com/developer/v3/endpoints/reports/#ebapi-get-reports-sales
func (c *Client) ReportSales(ctx context.Context, req *ReportRequest) (interface{}, error) {
	var v interface{}

	return v, c.getJSON(ctx, "/reports/sales/", req, &v)
}

// ReportSales returns a response of the aggregate attendees data
//
// https://www.eventbrite.com/developer/v3/endpoints/reports/#ebapi-get-reports-attendees
func (c *Client) ReportAttendees(ctx context.Context, req *ReportAttendees) (interface{}, error) {
	var v interface{}

	return v, c.getJSON(ctx, "/reports/attendees/", req, &v)
}
