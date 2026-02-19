package shippingproduct

type ShippingProductResponseItem struct {
	Name                     string                   `json:"name"`
	Carrier                  string                   `json:"carrier"`
	ServicePointsCarrier     string                   `json:"service_points_carrier"`
	AvailableFunctionalities AvailableFunctionalities `json:"available_functionalities"`
	Methods                  []Methods                `json:"methods"`
	Code                     string                   `json:"code"`
	WeightRange              WeightRange              `json:"weight_range"`
}
type AvailableFunctionalities struct {
	AgeCheck             []interface{} `json:"age_check"`
	B2B                  []bool        `json:"b2b"`
	B2C                  []bool        `json:"b2c"`
	Boxable              []bool        `json:"boxable"`
	BulkyGoods           []bool        `json:"bulky_goods"`
	CarrierBillingType   []interface{} `json:"carrier_billing_type"`
	CashOnDelivery       []interface{} `json:"cash_on_delivery"`
	DangerousGoods       []bool        `json:"dangerous_goods"`
	DeliveryAttempts     []interface{} `json:"delivery_attempts"`
	DeliveryBefore       []interface{} `json:"delivery_before"`
	DeliveryDeadline     []string      `json:"delivery_deadline"`
	DirectContractOnly   []bool        `json:"direct_contract_only"`
	EcoDelivery          []bool        `json:"eco_delivery"`
	FirstMile            []string      `json:"first_mile"`
	FlexDelivery         []bool        `json:"flex_delivery"`
	FormFactor           []string      `json:"form_factor"`
	FragileGoods         []bool        `json:"fragile_goods"`
	FreshGoods           []bool        `json:"fresh_goods"`
	HarmonizedLabel      []bool        `json:"harmonized_label"`
	IDCheck              []bool        `json:"id_check"`
	Incoterm             []interface{} `json:"incoterm"`
	Insurance            []interface{} `json:"insurance"`
	LastMile             []string      `json:"last_mile"`
	Manually             []bool        `json:"manually"`
	Multicollo           []bool        `json:"multicollo"`
	NeighborDelivery     []bool        `json:"neighbor_delivery"`
	NonConveyable        []bool        `json:"non_conveyable"`
	PersonalizedDelivery []bool        `json:"personalized_delivery"`
	Premium              []bool        `json:"premium"`
	Priority             []interface{} `json:"priority"`
	RegisteredDelivery   []bool        `json:"registered_delivery"`
	Returns              []bool        `json:"returns"`
	Segment              []interface{} `json:"segment"`
	ServiceArea          []string      `json:"service_area"`
	Signature            []bool        `json:"signature"`
	Size                 []interface{} `json:"size"`
	Sorted               []bool        `json:"sorted"`
	Surcharge            []bool        `json:"surcharge"`
	Tracked              []bool        `json:"tracked"`
	Tyres                []bool        `json:"tyres"`
	WeekendDelivery      []interface{} `json:"weekend_delivery"`
	Labelless            []bool        `json:"labelless"`
	Ers                  []bool        `json:"ers"`
}
type Functionalities struct {
}
type MaxDimensions struct {
	Length int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Properties struct {
	MinWeight     int           `json:"min_weight"`
	MaxWeight     int           `json:"max_weight"`
	MaxDimensions MaxDimensions `json:"max_dimensions"`
}
type Nl struct {
	Nl int `json:"NL"`
}
type LeadTimeHours struct {
	Nl Nl `json:"NL"`
}
type Methods struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name"`
	Functionalities     Functionalities `json:"functionalities"`
	ShippingProductCode string          `json:"shipping_product_code"`
	Properties          Properties      `json:"properties"`
	LeadTimeHours       LeadTimeHours   `json:"lead_time_hours"`
}
type WeightRange struct {
	MinWeight int `json:"min_weight"`
	MaxWeight int `json:"max_weight"`
}
