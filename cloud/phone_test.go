package cloud_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/familypulsa-dev/wapi/types"
)

func TestRegisterPhone(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	err := c.RegisterPhone(context.Background(), "123", "123456")
	if err != nil {
		t.Fatalf("RegisterPhone failed: %v", err)
	}
}

func TestCreatePhoneNumber(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	req := &types.CreatePhoneNumberRequest{
		CC:           "1",
		PhoneNumber:  "15551234",
		VerifiedName: "Lucky Shrub",
	}
	resp, err := c.CreatePhoneNumber(context.Background(), "waba-456", req)
	if err != nil {
		t.Fatalf("CreatePhoneNumber failed: %v", err)
	}
	if resp.ID != "110200345501442" {
		t.Errorf("expected 110200345501442, got %s", resp.ID)
	}
}

func TestRequestVerificationCode(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	err := c.RequestVerificationCode(context.Background(), "123", "SMS", "en_US")
	if err != nil {
		t.Fatalf("RequestVerificationCode failed: %v", err)
	}
}

func TestVerifyCode(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	err := c.VerifyCode(context.Background(), "123", "123830")
	if err != nil {
		t.Fatalf("VerifyCode failed: %v", err)
	}
}

func TestDeregisterPhone(t *testing.T) {
	ms := newMockServer()
	defer ms.Close()

	deregistered := false
	ms.on("POST", "/123/deregister", func(w http.ResponseWriter, r *http.Request) {
		deregistered = true
		writeJSON(w, http.StatusOK, map[string]bool{"success": true})
	})

	c := ms.client()
	err := c.DeregisterPhone(context.Background(), "123")
	if err != nil {
		t.Fatalf("DeregisterPhone failed: %v", err)
	}
	if !deregistered {
		t.Error("expected deregister to be called")
	}
}

func TestGetPhoneNumber(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	pn, err := c.GetPhoneNumber(context.Background(), "123")
	if err != nil {
		t.Fatalf("GetPhoneNumber failed: %v", err)
	}
	if pn.DisplayPhoneNumber != "+16505555555" {
		t.Errorf("expected +16505555555, got %s", pn.DisplayPhoneNumber)
	}
	if pn.QualityRating != "GREEN" {
		t.Errorf("expected GREEN, got %s", pn.QualityRating)
	}
}

func TestListPhoneNumbers(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	result, err := c.ListPhoneNumbers(context.Background(), "waba-456")
	if err != nil {
		t.Fatalf("ListPhoneNumbers failed: %v", err)
	}
	if result == nil || len(result.Data) == 0 {
		t.Fatal("expected at least one phone number")
	}
	if result.Data[0].DisplayPhoneNumber != "+16505555555" {
		t.Errorf("expected +16505555555, got %s", result.Data[0].DisplayPhoneNumber)
	}
	if result.Paging == nil {
		t.Fatal("expected paging in response")
	}
}

func TestSetTwoStepPIN(t *testing.T) {
	ms := newMockServer()
	defer ms.Close()

	var reqBody map[string]string
	ms.on("POST", "/123", func(w http.ResponseWriter, r *http.Request) {
		_ = parseBody(r, &reqBody)
		writeJSON(w, http.StatusOK, map[string]bool{"success": true})
	})

	c := ms.client()
	err := c.SetTwoStepPIN(context.Background(), "123", "654321")
	if err != nil {
		t.Fatalf("SetTwoStepPIN failed: %v", err)
	}
	if reqBody["pin"] != "654321" {
		t.Errorf("expected pin 654321, got %s", reqBody["pin"])
	}
}

func TestGetBusinessProfile(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	profile, err := c.GetBusinessProfile(context.Background(), "123")
	if err != nil {
		t.Fatalf("GetBusinessProfile failed: %v", err)
	}
	if profile.Description != "Test business profile" {
		t.Errorf("expected test description, got %s", profile.Description)
	}
}

func TestUpdateBusinessProfile(t *testing.T) {
	ms := newDefaultMockServer()
	defer ms.Close()

	c := ms.client()
	profile := &types.BusinessProfile{Description: "Updated description"}
	err := c.UpdateBusinessProfile(context.Background(), "123", profile)
	if err != nil {
		t.Fatalf("UpdateBusinessProfile failed: %v", err)
	}
}
