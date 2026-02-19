package sendcloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type WebhookAction string

const (
	WebhookActionIntegrationConnected WebhookAction = "integration_connected"
	WebhookActionIntegrationDeleted   WebhookAction = "integration_deleted"
	WebhookActionIntegrationUpdated   WebhookAction = "integration_updated"
	WebhookActionParcelStatusChanged  WebhookAction = "parcel_status_changed"
	WebhookActionReturnCreated        WebhookAction = "return_created"
)

type ParcelStatusChangedWebhookPayload struct {
	Action                       WebhookAction  `json:"action"`
	Timestamp                    int            `json:"timestamp"`
	CarrierStatusChangeTimestamp int            `json:"carrier_status_change_timestamp"`
	Parcel                       *WebhookParcel `json:"parcel"`
}

type Label struct {
	NormalPrinter []string `json:"normal_printer"`
	LabelPrinter  string   `json:"label_printer"`
}
type CustomsDeclaration struct {
}

type Data struct {
}
type WebhookCountry struct {
	Iso3 string `json:"iso_3"`
	Iso2 string `json:"iso_2"`
	Name string `json:"name"`
}

type WebhookParcel struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	CompanyName        string             `json:"company_name"`
	Address            string             `json:"address"`
	AddressDivided     AddressDivided     `json:"address_divided"`
	City               string             `json:"city"`
	PostalCode         string             `json:"postal_code"`
	Telephone          string             `json:"telephone"`
	Email              string             `json:"email"`
	DateCreated        string             `json:"date_created"`
	TrackingNumber     string             `json:"tracking_number"`
	Weight             string             `json:"weight"`
	Label              Label              `json:"label"`
	CustomsDeclaration CustomsDeclaration `json:"customs_declaration"`
	Status             Status             `json:"status"`
	Data               Data               `json:"data"`
	Country            WebhookCountry     `json:"country"`
	Shipment           Shipment           `json:"shipment"`
	OrderNumber        string             `json:"order_number"`
	ShipmentUUID       string             `json:"shipment_uuid"`
	ExternalOrderID    string             `json:"external_order_id"`
	ExternalShipmentID string             `json:"external_shipment_id"`
}

const signatureHeader = "Sendcloud-Signature"
const maxBodySize = 1024 * 1024 * 5 // 5 MiB

func ParseWebhook(secretKey string, h http.Header, r *http.Request) (*ParcelStatusChangedWebhookPayload, error) {
	// Read up to 5 MiB
	limitedReader := io.LimitReader(r.Body, maxBodySize)
	body, err := io.ReadAll(limitedReader)
	if err != nil {
		return nil, err
	}

	// Compute HMAC SHA256
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, _ = mac.Write(body)
	expectedSig := hex.EncodeToString(mac.Sum(nil))

	// Get provided signature
	providedSig := h.Get(signatureHeader)
	if providedSig == "" {
		return nil, errors.New("missing signature header")
	}

	// Constant-time compare
	if !hmac.Equal([]byte(expectedSig), []byte(providedSig)) {
		return nil, errors.New("invalid signature")
	}

	// Unmarshal payload
	var payload ParcelStatusChangedWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}

	if payload.Action != WebhookActionParcelStatusChanged {
		return nil, errors.New("decoding action type not supported")
	}

	return &payload, nil
}
