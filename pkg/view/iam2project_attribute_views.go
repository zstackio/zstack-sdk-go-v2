// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectAttributeInventoryView IAM2ProjectAttribute
type IAM2ProjectAttributeInventoryView struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryIAM2ProjectAttributeView QueryIAM2ProjectAttribute
type QueryIAM2ProjectAttributeView struct {
	Inventories []IAM2ProjectAttributeInventoryView `json:"inventories,omitempty"`
}

// UpdateIAM2ProjectAttributeEventView UpdateIAM2ProjectAttributeEvent
type UpdateIAM2ProjectAttributeEventView struct {
	Inventory IAM2ProjectAttributeInventoryView `json:"inventory,omitempty"`
}

