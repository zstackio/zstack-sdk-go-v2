// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IAM2VirtualIDGroupAttributeInventoryView IAM2VirtualIDGroupAttribute
type IAM2VirtualIDGroupAttributeInventoryView struct {
	BaseInfoView
	BaseTimeView
	GroupUuid string `json:"groupUuid,omitempty"`
	Value string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateIAM2VirtualIDGroupAttributeEventView UpdateIAM2VirtualIDGroupAttributeEvent
type UpdateIAM2VirtualIDGroupAttributeEventView struct {
	Inventory IAM2VirtualIDGroupAttributeInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2VirtualIDGroupAttributeView QueryIAM2VirtualIDGroupAttribute
type QueryIAM2VirtualIDGroupAttributeView struct {
	Inventories []IAM2VirtualIDGroupAttributeInventoryView `json:"inventories,omitempty"`
}

