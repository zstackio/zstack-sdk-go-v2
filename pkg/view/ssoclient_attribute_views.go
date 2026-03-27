// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SSOClientAttributeInventoryView SSOClientAttribute
type SSOClientAttributeInventoryView struct {
	BaseInfoView
	BaseTimeView
	Value string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	SsoClientUuid string `json:"ssoClientUuid,omitempty"`
}

// UpdateSSOClientAttributeEventView UpdateSSOClientAttributeEvent
type UpdateSSOClientAttributeEventView struct {
	Inventory SSOClientAttributeInventoryView `json:"inventory,omitempty"`
}

