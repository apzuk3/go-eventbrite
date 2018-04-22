package eventbrite

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"

	"gopkg.in/go-playground/validator.v9"
)

var (
	defaultRequestsPerSecond = 5
	validate                 = validator.New()
)

// Client may be used to make requests to the Eventbrite API
type Client struct {
	httpClient        *http.Client
	token             string
	baseURL           string
	requestsPerSecond int
	ratePerSecond     chan int
}

// ClientOption is the type of constructor options for NewClient(...).
type ClientOption func(*Client) error

// NewClient constructs a new Client which can make requests to the Eventbrite API.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{}

	WithBaseURL("https://www.eventbriteapi.com/v3")(c)
	WithRateLimit(defaultRequestsPerSecond)(c)
	WithHTTPClient(&http.Client{})(c)

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	if c.requestsPerSecond > 0 {
		c.ratePerSecond = make(chan int, c.requestsPerSecond)
		for i := 0; i < c.requestsPerSecond; i++ {
			c.ratePerSecond <- 1
		}
		go func() {
			time.Sleep(time.Second)
			for range time.Tick(time.Second / time.Duration(c.requestsPerSecond)) {
				c.ratePerSecond <- 1
			}
		}()
	}

	return c, nil
}

// WithHTTPClient configures a Eventbrite client with a http.Client to make requests over.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// WithPersonalToken configures a Eventbrite API client with auth token
func WithToken(token string) ClientOption {
	return func(c *Client) error {
		c.token = token
		return nil
	}
}

// WithBaseURL configures a Eventbrite API client with a custom base url
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// WithRateLimit configures the rate limit for back end requests. Default is to
// limit to 50 requests per second. A value of zero disables rate limiting.
func WithRateLimit(requestsPerSecond int) ClientOption {
	return func(c *Client) error {
		c.requestsPerSecond = requestsPerSecond
		return nil
	}
}

func (c *Client) awaitRateLimiter(ctx context.Context) error {
	if c.ratePerSecond == nil {
		return nil
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c.ratePerSecond:
		return nil
	}
}

func (c *Client) get(ctx context.Context, path string, apiReq interface{}) (*http.Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	if apiReq != nil {
		if err := validate.Struct(apiReq); err != nil {
			return nil, err
		}
	}

	host := path
	if c.baseURL != "" {
		host = c.baseURL
	}
	req, err := http.NewRequest(http.MethodGet, host+path, nil)
	if err != nil {
		return nil, err
	}
	q, err := c.generateAuthQuery(path, toValues(apiReq))
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = q
	return ctxhttp.Do(ctx, c.httpClient, req)
}

func (c *Client) delete(ctx context.Context, path string) (*http.Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	host := path
	if c.baseURL != "" {
		host = c.baseURL
	}

	req, err := http.NewRequest(http.MethodDelete, host+path, nil)
	if err != nil {
		return nil, err
	}

	q, err := c.generateAuthQuery(path, url.Values{})
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = q
	return ctxhttp.Do(ctx, c.httpClient, req)
}

func (c *Client) post(ctx context.Context, path string, apiReq interface{}) (*http.Response, error) {

	if apiReq != nil {
		if err := validate.Struct(apiReq); err != nil {
			return nil, err
		}
	}

	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	host := path
	if c.baseURL != "" {
		host = c.baseURL
	}

	body, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, host+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q, err := c.generateAuthQuery(path, url.Values{})
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = q
	return ctxhttp.Do(ctx, c.httpClient, req)
}

func (c *Client) generateAuthQuery(path string, q url.Values) (string, error) {
	if c.token != "" {
		q.Set("token", c.token)
		q.Set("expand", "venue,category,subcategories")
		return q.Encode(), nil
	}
	return "", errors.New("eventbrite: Token missing")
}

func (c *Client) getJSON(ctx context.Context, path string, apiReq interface{}, resp interface{}) error {
	httpResp, err := c.get(ctx, path, apiReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode == 200 {
		return json.NewDecoder(httpResp.Body).Decode(resp)
	}

	respErr := Error{}
	json.NewDecoder(httpResp.Body).Decode(&respErr)
	return respErr
}

func (c *Client) postJSON(ctx context.Context, path string, apiReq interface{}, resp interface{}) error {
	httpResp, err := c.post(ctx, path, apiReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func (c *Client) deleteJSON(ctx context.Context, path string, resp interface{}) error {
	httpResp, err := c.delete(ctx, path)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func toValues(i interface{}) (values url.Values) {

	if i == nil {
		return url.Values{}
	}

	switch i.(type) {
	case url.Values:
		return i.(url.Values)
	}

	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		tag := typ.Field(i).Tag.Get("json")
		if tag == "" {
			tag = typ.Field(i).Name
		}
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}

		if v != "" {
			values.Set(tag, v)
		}
	}
	return
}
