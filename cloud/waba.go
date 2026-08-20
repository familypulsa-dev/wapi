package cloud

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	wapi "github.com/familypulsa-dev/wapi"
	"github.com/familypulsa-dev/wapi/types"
)

// ListWhatsAppBusinessAccounts returns all WABAs owned by the given business ID, with optional pagination.
func (c *CloudClient) ListWhatsAppBusinessAccounts(ctx context.Context, businessID string, opts ...wapi.ListOption) (*types.WABAList, error) {
	params := &wapi.ListParams{}
	for _, opt := range opts {
		opt(params)
	}

	v := url.Values{}
	v.Set("fields", "id,name,currency,timezone_id,message_template_namespace,status,whatsapp_business_manager_messaging_limit")
	if params.Limit > 0 {
		v.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.Offset != "" {
		v.Set("after", params.Offset)
	}

	path := fmt.Sprintf("%s/owned_whatsapp_business_accounts", businessID)
	var list types.WABAList
	if err := c.doGet(ctx, path, v, &list); err != nil {
		return nil, fmt.Errorf("list waba accounts: %w", err)
	}
	return &list, nil
}

// GetWABAPricingAnalytics fetches pricing analytics for a WABA.
func (c *CloudClient) GetWABAPricingAnalytics(ctx context.Context, wabaID string, start, end int64) (*types.WhatsAppPricingAnalyticsResponse, error) {
	// e.g. fields=pricing_analytics.start({start}).end({end}).granularity(DAILY).phone_numbers([]).dimensions(["PHONE","PRICING_CATEGORY"])
	v := url.Values{}
	fields := fmt.Sprintf(`pricing_analytics.start(%d).end(%d).granularity(DAILY).phone_numbers([]).dimensions(["PHONE","PRICING_CATEGORY"])`, start, end)
	v.Set("fields", fields)

	var res types.WhatsAppPricingAnalyticsResponse
	if err := c.doGet(ctx, wabaID, v, &res); err != nil {
		return nil, fmt.Errorf("get waba pricing analytics: %w", err)
	}
	return &res, nil
}
