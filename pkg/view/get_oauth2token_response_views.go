// Copyright (c) ZStack.io, Inc.

package view

// GetOAuth2TokenView GetOAuth2Token
type GetOAuth2TokenView struct {
	Inventory OAuth2TokenInventoryView `json:"inventory,omitempty"`
	ServerTokenInventory SSOServerTokenInventoryView `json:"serverTokenInventory,omitempty"`
	AdditionalTokenInventory map[string]interface{} `json:"additionalTokenInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

