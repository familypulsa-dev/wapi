package types

// WhatsAppBusinessAccount represents a WhatsApp Business Account (WABA) returned by owned_whatsapp_business_accounts.
type WhatsAppBusinessAccount struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	Currency                 string `json:"currency,omitempty"`
	TimezoneID               string `json:"timezone_id,omitempty"`
	MessageTemplateNamespace string `json:"message_template_namespace,omitempty"`
}

// WABAList is the paginated response from listing owned WABAs.
type WABAList struct {
	Data   []*WhatsAppBusinessAccount `json:"data"`
	Paging *Paging                    `json:"paging,omitempty"`
}

// WhatsAppPricingAnalyticsResponse represents the response from the pricing_analytics endpoint.
type WhatsAppPricingAnalyticsResponse struct {
	PricingAnalytics *PricingAnalyticsData `json:"pricing_analytics,omitempty"`
	ID               string                `json:"id,omitempty"`
}

type PricingAnalyticsData struct {
	Data []PricingAnalyticsGroup `json:"data,omitempty"`
}

type PricingAnalyticsGroup struct {
	DataPoints []PricingDataPoint `json:"data_points,omitempty"`
}

type PricingDataPoint struct {
	Start           int64   `json:"start"`
	End             int64   `json:"end"`
	PhoneNumber     string  `json:"phone_number"`
	PricingCategory string  `json:"pricing_category"`
	Volume          int     `json:"volume"`
	Cost            float64 `json:"cost"`
}
