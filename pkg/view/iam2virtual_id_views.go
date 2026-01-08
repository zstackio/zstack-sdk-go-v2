// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDInventoryView IAM2VirtualID
type IAM2VirtualIDInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type       string                       `json:"type,omitempty"`
	State      string                       `json:"state,omitempty"`
	Attributes []IAM2AttributeInventoryView `json:"attributes,omitempty"`
}

// QueryIAM2VirtualIDView QueryIAM2VirtualID
type QueryIAM2VirtualIDView struct {
	Inventories []IAM2VirtualIDInventoryView `json:"inventories,omitempty"`
}

// CreateIAM2VirtualIDEventView CreateIAM2VirtualIDEvent
type CreateIAM2VirtualIDEventView struct {
	Inventory IAM2VirtualIDInventoryView `json:"inventory,omitempty"`
}

// ChangeIAM2VirtualIDTypeEventView ChangeIAM2VirtualIDTypeEvent
type ChangeIAM2VirtualIDTypeEventView struct {
	Inventory IAM2VirtualIDInventoryView `json:"inventory,omitempty"`
}

// CleanInvalidLdapIAM2BindingEventView CleanInvalidLdapIAM2BindingEvent
type CleanInvalidLdapIAM2BindingEventView struct {
	Inventories []IAM2VirtualIDInventoryView `json:"inventories,omitempty"`
}

// GetIAM2VirtualIDInGroupView GetIAM2VirtualIDInGroup
type GetIAM2VirtualIDInGroupView struct {
	Inventories []IAM2VirtualIDInventoryView `json:"inventories,omitempty"`
	Total       int64                        `json:"total,omitempty"`
	Success     bool                         `json:"success,omitempty"`
}

// DeleteIAM2VirtualIDEventView DeleteIAM2VirtualIDEvent
type DeleteIAM2VirtualIDEventView struct {
	Success bool `json:"success,omitempty"`
}

// LoginIAM2VirtualIDView LoginIAM2VirtualID
type LoginIAM2VirtualIDView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// UpdateIAM2VirtualIDEventView UpdateIAM2VirtualIDEvent
type UpdateIAM2VirtualIDEventView struct {
	Inventory IAM2VirtualIDInventoryView `json:"inventory,omitempty"`
}

// ChangeIAM2VirtualIDStateEventView ChangeIAM2VirtualIDStateEvent
type ChangeIAM2VirtualIDStateEventView struct {
	Inventory IAM2VirtualIDInventoryView `json:"inventory,omitempty"`
}
