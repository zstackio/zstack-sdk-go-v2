// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SSOClientAttributeInventoryView SSOClientAttribute
type SSOClientAttributeInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	SsoClientUuid *string `json:"ssoClientUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateSSOClientAttributeEventView UpdateSSOClientAttributeEvent
type UpdateSSOClientAttributeEventView struct {
	Inventory SSOClientAttributeInventoryView `json:"inventory,omitempty"`
}

