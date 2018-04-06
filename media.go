package eventbrite

import (
	"fmt"
	"golang.org/x/net/context"
)

// https://www.eventbrite.com/developer/v3/resources/uploads/
type Media struct {
	// the method (always POST)
	Method string `json:"upload_method"`
	// oauth token
	Token string `json:"upload_token"`
	// the URL that should be uploaded to
	Url string `json:"url"`
	// the POST data that should be included in the POST that uploads the file
	UploadData UploadData `json:"upload_data"`
	// Specifies the POST field that the file itself should be included in (handled using HTTP multipart upload)
	FileParameterName string `json:"file_parameter_name"`
}

type UploadData struct {
	AWSAccessKeyID string `json:"AWSAccessKeyId"`
	Bucket         string `json:"bucket"`
	Acl            string `json:"acl"`
	Key            string `json:"kcl"`
	Signature      string `json:"signature"`
	Policy         string `json:"policy"`
}

// https://www.eventbrite.com/developer/v3/endpoints/media/#ebapi-id1
type MediaGetUpload struct {
	// The type of image to upload (Valid choices are: image-event-logo, image-event-logo-preserve-quality,
	// image-event-view-from-seat, image-organizer-logo, image-user-photo, or image-structured-content)
	Type string `json:"type" validate:"required"`
}

// https://www.eventbrite.com/developer/v3/endpoints/media/#ebapi-id3
type MediaCreateUpload struct {
	// The upload_token from the GET portion of the upload
	UploadToken string `json:"upload_token" validate:"required"`
	// X coordinate for top-left corner of crop mask
	TopLeftX int `json:"crop_mask.top_left.x"`
	// Y coordinate for top-left corner of crop mask
	TopLeftY int `json:"crop_mask.top_left.y"`
	// Crop mask width
	Width int `json:"crop_mask.width"`
	// Crop mask height
	Height int `json:"crop_mask.height"`
}

// https://www.eventbrite.com/developer/v3/endpoints/media/#ebapi-get-media-upload
// https://www.eventbrite.com/developer/v3/resources/uploads/
func (c *Client) MediaGet(ctx context.Context, req *MediaGetUpload) (*Media, error) {
	m := new(Media)

	return m, c.getJSON(ctx, "/media/upload/", req, m)
}

// Return an image for a given id
//
// https://www.eventbrite.com/developer/v3/endpoints/media/#ebapi-get-media-id
func (c *Client) MediaGetUpload(ctx context.Context, id string) (*Image, error) {
	i := new(Image)

	return i, c.getJSON(ctx, fmt.Sprintf("/media/%s/", id), nil, i)
}

// https://www.eventbrite.com/developer/v3/endpoints/media/#ebapi-post-media-upload
func (c *Client) MediaCreate(ctx context.Context, req *MediaCreateUpload) (*Image, error) {
	i := new(Image)

	return i, c.getJSON(ctx, fmt.Sprintf("/media/upload/"), req, i)
}
