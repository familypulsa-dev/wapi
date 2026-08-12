package cloud

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/familypulsa-dev/wapi"
	"github.com/familypulsa-dev/wapi/types"
)

// CreatePhoneNumber creates a new business phone number on a WABA.
func (c *CloudClient) CreatePhoneNumber(ctx context.Context, wabaID string, req *types.CreatePhoneNumberRequest) (*types.CreatePhoneNumberResponse, error) {
	path := fmt.Sprintf("%s/phone_numbers", wabaID)
	var resp types.CreatePhoneNumberResponse
	if err := c.do(ctx, "POST", path, req, &resp); err != nil {
		return nil, fmt.Errorf("create phone number: %w", err)
	}
	return &resp, nil
}

// UpdateDisplayName requests a display name change for a phone number.
// The new name is subject to Meta review via name_status.
func (c *CloudClient) UpdateDisplayName(ctx context.Context, phoneNumberID string, req *types.UpdateDisplayNameRequest) error {
	if req.MessagingProduct == "" {
		req.MessagingProduct = "whatsapp"
	}
	path := fmt.Sprintf("%s", phoneNumberID)
	return c.do(ctx, "POST", path, req, nil)
}

// RequestVerificationCode sends a verification code via SMS or VOICE.
func (c *CloudClient) RequestVerificationCode(ctx context.Context, phoneNumberID, codeMethod, language string) error {
	path := fmt.Sprintf("%s/request_code?code_method=%s&language=%s", phoneNumberID, codeMethod, language)
	return c.do(ctx, "POST", path, nil, nil)
}

// VerifyCode confirms the verification code for a phone number.
func (c *CloudClient) VerifyCode(ctx context.Context, phoneNumberID, code string) error {
	path := fmt.Sprintf("%s/verify_code?code=%s", phoneNumberID, code)
	return c.do(ctx, "POST", path, nil, nil)
}

// RegisterPhone registers a phone number with a 6-digit PIN.
func (c *CloudClient) RegisterPhone(ctx context.Context, phoneNumberID, pin string) error {
	body := map[string]string{
		"messaging_product": "whatsapp",
		"pin":               pin,
	}
	path := fmt.Sprintf("%s/register", phoneNumberID)
	return c.do(ctx, "POST", path, body, nil)
}

// DeregisterPhone deregisters a phone number from the WABA.
func (c *CloudClient) DeregisterPhone(ctx context.Context, phoneNumberID string) error {
	path := fmt.Sprintf("%s/deregister", phoneNumberID)
	return c.do(ctx, "POST", path, map[string]string{"messaging_product": "whatsapp"}, nil)
}

const phoneFields = "id,display_phone_number,verified_name,quality_rating,name_status,code_verification_status,status,pin_enabled"

// GetPhoneNumber returns details for a specific phone number.
func (c *CloudClient) GetPhoneNumber(ctx context.Context, phoneNumberID string) (*types.PhoneNumber, error) {
	path := fmt.Sprintf("%s?fields=%s", phoneNumberID, phoneFields)
	var pn types.PhoneNumber
	if err := c.do(ctx, "GET", path, nil, &pn); err != nil {
		return nil, fmt.Errorf("get phone number: %w", err)
	}
	return &pn, nil
}

// ListPhoneNumbers returns phone numbers associated with a WABA with optional pagination
// (wapi.WithLimit, wapi.WithOffset).
func (c *CloudClient) ListPhoneNumbers(ctx context.Context, wabaID string, opts ...wapi.ListOption) (*types.PhoneNumberList, error) {
	params := &wapi.ListParams{}
	for _, opt := range opts {
		opt(params)
	}

	path := fmt.Sprintf("%s/phone_numbers?fields=%s", wabaID, phoneFields)
	if params.Limit > 0 {
		path += "&limit=" + strconv.Itoa(params.Limit)
	}
	if params.Offset != "" {
		path += "&after=" + url.QueryEscape(params.Offset)
	}

	var result types.PhoneNumberList
	if err := c.do(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("list phone numbers: %w", err)
	}
	return &result, nil
}

// SetTwoStepPIN enables or changes the 6-digit PIN for two-step verification.
func (c *CloudClient) SetTwoStepPIN(ctx context.Context, phoneNumberID, pin string) error {
	body := map[string]string{"pin": pin}
	return c.do(ctx, "POST", phoneNumberID, body, nil)
}
