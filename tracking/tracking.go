package tracking

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/afosto/sendcloud-go"
)

type TrackingTime struct {
	time.Time
}

func (d *TrackingTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	// "2022-06-16 15:04:42.663125+00:00",
	t, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", s)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

type ShortTime struct {
	time.Time
}

func (d *ShortTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	// "2022-06-16",
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

type TrackingUpdateResponse struct {
	CarrierUpdateTimestamp TrackingTime `json:"carrier_update_timestamp"`
	ParcelStatusHistoryID  string       `json:"parcel_status_history_id"`
	ParentStatus           string       `json:"parent_status"`
	CarrierCode            string       `json:"carrier_code"`
	CarrierMessage         string       `json:"carrier_message"`
}

type TrackingResponse struct {
	ParcelID             string                   `json:"parcel_id"`
	CarrierCode          string                   `json:"carrier_code"`
	CreatedAt            TrackingTime             `json:"created_at"`
	CarrierTrackingURL   string                   `json:"carrier_tracking_url"`
	SendcloudTrackingURL string                   `json:"sendcloud_tracking_url"`
	IsReturn             bool                     `json:"is_return"`
	IsToServicePoint     bool                     `json:"is_to_service_point"`
	IsMailBox            bool                     `json:"is_mail_box"`
	ExpectedDeliveryDate ShortTime                `json:"expected_delivery_date"`
	Statuses             []TrackingUpdateResponse `json:"statuses"`
}

func (s TrackingResponse) GetResponse() interface{} {
	return s
}

func (s *TrackingResponse) SetResponse(body []byte) error {
	err := json.Unmarshal(body, &s)
	if err != nil {
		return err
	}
	return nil
}

type Client struct {
	apiKey    string
	apiSecret string
}

func New(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (service Client) GetTracking(trackingNumber string) (*TrackingResponse, error) {
	uri, _ := url.Parse("https://panel.sendcloud.sc/api/v2/tracking/" + trackingNumber)

	var resp TrackingResponse
	if err := sendcloud.Request("GET", uri.String(), nil, service.apiKey, service.apiSecret, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
