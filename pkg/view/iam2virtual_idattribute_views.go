// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDAttributeInventoryView IAM2VirtualIDAttribute
type IAM2VirtualIDAttributeInventoryView struct {
	VirtualIDUuid *string `json:"virtualIDUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateIAM2VirtualIDAttributeEventView UpdateIAM2VirtualIDAttributeEvent
type UpdateIAM2VirtualIDAttributeEventView struct {
	Inventory IAM2VirtualIDAttributeInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2VirtualIDAttributeView QueryIAM2VirtualIDAttribute
type QueryIAM2VirtualIDAttributeView struct {
	Inventories []IAM2VirtualIDAttributeInventoryView `json:"inventories,omitempty"`
}

