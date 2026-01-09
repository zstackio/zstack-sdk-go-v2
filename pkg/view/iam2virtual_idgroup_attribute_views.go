// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDGroupAttributeInventoryView IAM2VirtualIDGroupAttribute
type IAM2VirtualIDGroupAttributeInventoryView struct {
	GroupUuid *string `json:"groupUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateIAM2VirtualIDGroupAttributeEventView UpdateIAM2VirtualIDGroupAttributeEvent
type UpdateIAM2VirtualIDGroupAttributeEventView struct {
	Inventory IAM2VirtualIDGroupAttributeInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2VirtualIDGroupAttributeView QueryIAM2VirtualIDGroupAttribute
type QueryIAM2VirtualIDGroupAttributeView struct {
	Inventories []IAM2VirtualIDGroupAttributeInventoryView `json:"inventories,omitempty"`
}

