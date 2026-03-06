// Copyright (c) ZStack.io, Inc.
// Auto-generated view tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

func TestSSOClientInventoryView_UnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"uuid": "test-uuid-001",
		"name": "test-ssoclient",
		"createDate": "2024-01-01T00:00:00.000+08:00",
		"lastOpDate": "2024-01-01T00:00:00.000+08:00"
	}`
	var v view.SSOClientInventoryView
	err := json.Unmarshal([]byte(jsonStr), &v)
	assertNoError(t, err)
}

func TestSSOClientInventoryView_UnmarshalEmpty(t *testing.T) {
	var v view.SSOClientInventoryView
	err := json.Unmarshal([]byte(`{}`), &v)
	assertNoError(t, err)
}

func TestSSOClientInventoryView_UnmarshalFull(t *testing.T) {
	jsonStr := `{
		"uuid": "test-uuid-001",
		"name": "test-ssoclient",
		"description": "test description",
		"clientType": "OIDC",
		"loginType": "sso",
		"loginMNUrl": "https://example.com/login",
		"redirectUrl": "https://example.com/callback",
		"accountUuid": "account-uuid-001",
		"attributes": [
			{"uuid": "attr-uuid-001", "name": "attr1"}
		],
		"createDate": "2024-01-01T00:00:00.000+08:00",
		"lastOpDate": "2024-01-01T00:00:00.000+08:00"
	}`
	var v view.SSOClientInventoryView
	err := json.Unmarshal([]byte(jsonStr), &v)
	assertNoError(t, err)
	assertEqual(t, "test-uuid-001", v.UUID)
	assertEqual(t, "OIDC", v.ClientType)
	assertEqual(t, "sso", v.LoginType)
	assertEqual(t, "https://example.com/login", v.LoginMNUrl)
	assertEqual(t, "https://example.com/callback", v.RedirectUrl)
	assertEqual(t, "account-uuid-001", v.AccountUuid)
	if len(v.Attributes) != 1 {
		t.Fatalf("expected 1 attribute, got %d", len(v.Attributes))
	}
}

func TestDeleteSSOClientEventView_UnmarshalJSON(t *testing.T) {
	jsonStr := `{"success": true}`
	var v view.DeleteSSOClientEventView
	err := json.Unmarshal([]byte(jsonStr), &v)
	assertNoError(t, err)
	if !v.Success {
		t.Fatal("expected success to be true")
	}
}

func TestDeleteSSOClientEventView_UnmarshalEmpty(t *testing.T) {
	var v view.DeleteSSOClientEventView
	err := json.Unmarshal([]byte(`{}`), &v)
	assertNoError(t, err)
}

func TestGetSSOClientView_UnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"inventories": [
			{
				"uuid": "sso-uuid-001",
				"name": "sso-client-1",
				"clientType": "OIDC",
				"loginType": "sso"
			},
			{
				"uuid": "sso-uuid-002",
				"name": "sso-client-2",
				"clientType": "SAML2",
				"loginType": "sso"
			}
		]
	}`
	var v view.GetSSOClientView
	err := json.Unmarshal([]byte(jsonStr), &v)
	assertNoError(t, err)
	if len(v.Inventories) != 2 {
		t.Fatalf("expected 2 inventories, got %d", len(v.Inventories))
	}
	assertEqual(t, "sso-uuid-001", v.Inventories[0].UUID)
	assertEqual(t, "OIDC", v.Inventories[0].ClientType)
	assertEqual(t, "sso-uuid-002", v.Inventories[1].UUID)
	assertEqual(t, "SAML2", v.Inventories[1].ClientType)
}

func TestGetSSOClientView_UnmarshalEmpty(t *testing.T) {
	var v view.GetSSOClientView
	err := json.Unmarshal([]byte(`{}`), &v)
	assertNoError(t, err)
	if len(v.Inventories) != 0 {
		t.Fatalf("expected 0 inventories, got %d", len(v.Inventories))
	}
}

