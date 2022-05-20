package mixpaysdk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	MixPayEndpoint       = "https://api.mixpay.me"
	PaymentsUri          = "/v1/payments"
	PaymentsEstimatedUri = "/v1/payments_estimated"
	PaymentsResultUri    = "/v1/payments_result"
)

type Client struct {
	client *resty.Client
}

func New() *Client {
	c := resty.New()
	c.SetBaseURL(MixPayEndpoint)
	return &Client{client: c}
}


func (c *Client) UnmarshalResponse(response *resty.Response, v any) error {
	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("%s %s", response.Status(), response.Body())
	}
	if err := json.Unmarshal(response.Body(), &v); err != nil {
		return fmt.Errorf("json unmarshal error %s", err.Error())
	}
	return nil
}