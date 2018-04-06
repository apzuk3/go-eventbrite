package eventbrite

import (
	"fmt"

	"golang.org/x/net/context"
)

// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-parameters
type CreateTrackingBeaconRequest struct {
	// The tracking pixel third party type. Allowed types are: Facebook Pixel, Twitter Ads,
	// AdWords, Google Analytics, Simple Image Pixel, Adroll iPixel
	TrackingType string `json:"tracking_type" validate:"required"`

	// The Event ID of the event that this tracking beacon will fire in
	EventID string `json:"event_id"`

	// The User ID wherein the tracking beacon will be assigned to all of this user’s events
	UserID string `json:"user_id"`

	// The Pixel ID given by the third party that will fire when a attendee lands on the page you are tracking
	PixelID string `json:"pixel_id"`

	// The additional pixel data needed to determine which page to fire the tracking pixel on
	Triggers interface{} `json:"triggers"`
}

// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-id3
type UpdateTrackingBeaconRequest struct {
	// The tracking pixel third party type. Allowed types are: Facebook Pixel, Twitter Ads,
	// AdWords, Google Analytics, Simple Image Pixel, Adroll iPixel
	TrackingType string `json:"tracking_type" validate:"required"`

	// The Event ID of the event that this tracking beacon will fire in
	EventID string `json:"event_id"`

	// The User ID wherein the tracking beacon will be assigned to all of this user’s events
	UserID string `json:"user_id"`

	// The Pixel ID given by the third party that will fire when a attendee lands on the page you are tracking
	PixelID string `json:"pixel_id"`

	// The additional pixel data needed to determine which page to fire the tracking pixel on
	Triggers interface{} `json:"triggers"`
}

// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-id1
type GetTrackingBeaconRequest struct {
	// returned format
	ReturnFmt string `json:"return_fmt"`
}

type GetTrackingBeaconForEventRequest struct {
	// returned format
	ReturnFmt string `json:"return_fmt"`
}

type GetTrackingBeaconForUserRequest struct {
	// returned format
	ReturnFmt string `json:"return_fmt"`
}

// TrackingBeaconCreate makes a new tracking beacon. Returns an tracking_beacon as tracking_beacon. Either event_id
// or user_id is required for each tracking beacon. If the event_id is provided, the tracking pixel will fire only for
// that event. If the user_id is provided, the tracking pixel will fire for all events organized by that user
//
// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-post-tracking-beacons
func (c *Client) TrackingBeaconCreate(ctx context.Context, req *CreateTrackingBeaconRequest) (*TrackingBeacon, error) {
	res := new(TrackingBeacon)

	return res, c.postJSON(ctx, "/tracking_beacons/", req, res)
}

// TrackingBeaconGet returns the tracking_beacon with the specified :tracking_beacons_id
//
// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-get-tracking-beacons-tracking-beacons-id
func (c *Client) TrackingBeaconGet(ctx context.Context, id string, req *GetTrackingBeaconRequest) (*TrackingBeacon, error) {
	res := new(TrackingBeacon)

	return res, c.getJSON(ctx, "/tracking_beacons/"+id, req, res)
}

// TrackingBeaconGet updates the tracking_beacons with the specified :tracking_beacons_id. Though event_id and
// user_id are not individually required, it is a requirement to have a tracking beacon where either one must exist.
// Returns an tracking_beacon as tracking_beacon
//
// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-post-tracking-beacons-tracking-beacons-id
func (c *Client) TrackingBeaconUpdate(ctx context.Context, id string, req *UpdateTrackingBeaconRequest) (*TrackingBeacon, error) {
	res := new(TrackingBeacon)

	return res, c.postJSON(ctx, "/tracking_beacons/"+id, req, res)
}

// TrackingBeaconDelete delete the tracking_beacons with the specified :tracking_beacons_id
//
// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-delete-tracking-beacons-tracking-beacons-id
func (c *Client) TrackingBeaconDelete(ctx context.Context, id string) (*TrackingBeacon, error) {
	res := new(TrackingBeacon)

	return res, c.deleteJSON(ctx, "/tracking_beacons/"+id, res)
}

// TrackingBeaconGetForEvent returns the list of tracking_beacon for the event :event_id
//
// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-get-events-event-id-tracking-beacons
func (c *Client) TrackingBeaconGetForEvent(ctx context.Context, eventId string, req *GetTrackingBeaconForEventRequest) (*TrackingBeacon, error) {
	res := new(TrackingBeacon)

	return res, c.getJSON(ctx, fmt.Sprintf("/events/%s/tracking_beacons/", eventId), req, res)
}

// TrackingBeaconGetForUser returns the list of tracking_beacon for the user :user_id
//
// https://www.eventbrite.com/developer/v3/endpoints/tracking_beacons/#ebapi-get-users-user-id-tracking-beacons
func (c *Client) TrackingBeaconGetForUser(ctx context.Context, userId string, req *GetTrackingBeaconForUserRequest) (*TrackingBeacon, error) {
	res := new(TrackingBeacon)

	return res, c.getJSON(ctx, fmt.Sprintf("/events/%s/tracking_beacons/", userId), req, res)
}
