package shippingproduct

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/afosto/sendcloud-go"
)

type WeightUnit string
type LengthUnit string

const (
	Gram     WeightUnit = "gram"
	Kilogram WeightUnit = "kilogram"

	Meter      LengthUnit = "meter"
	Centimeter LengthUnit = "centimeter"
	Millimeter LengthUnit = "millimeter"
)

type ShippingPriceParams struct {
	ShippingMethodId int

	FromCountry string // FromCountry is an Alpha2 country code, e.g. SE
	ToCountry   string // ToCountry is an Alpha2 country code, e.g. SE

	Weight     int
	WeightUnit WeightUnit

	Length     int
	LengthUnit LengthUnit
	Width      int
	WidthUnit  LengthUnit
	Height     int
	HeightUnit LengthUnit

	LeadTimeHours int
}

type ShippingProductResponse []ShippingProductResponseItem

func (s ShippingProductResponse) GetResponse() interface{} {
	return s
}

func (s *ShippingProductResponse) SetResponse(body []byte) error {
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

// Returns the sendcloud pickup point ID mapped from a SPID ID
func (service Client) GetShippingPrice(params ShippingPriceParams) (ShippingProductResponse, error) {
	//prepare bounding box url
	uri, _ := url.Parse("https://panel.sendcloud.sc/api/v2/shipping-products/")
	paramsContainer := uri.Query()
	paramsContainer.Set("from_country", params.FromCountry)
	paramsContainer.Set("to_country", params.ToCountry)
	paramsContainer.Set("weight", fmt.Sprintf("%d", params.Weight))
	paramsContainer.Set("weight_unit", string(params.WeightUnit))

	uri.RawQuery = paramsContainer.Encode()

	var prices ShippingProductResponse
	if err := sendcloud.Request("GET", uri.String(), nil, service.apiKey, service.apiSecret, &prices); err != nil {
		return nil, err
	}

	return prices, nil
}
