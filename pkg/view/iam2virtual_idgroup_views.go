// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDGroupInventoryView IAM2VirtualIDGroup
type IAM2VirtualIDGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	ProjectUuid string                       `json:"projectUuid,omitempty"`
	State       string                       `json:"state,omitempty"`
	Attributes  []IAM2AttributeInventoryView `json:"attributes,omitempty"`
}

// QueryIAM2VirtualIDGroupView QueryIAM2VirtualIDGroup
type QueryIAM2VirtualIDGroupView struct {
	Inventories []IAM2VirtualIDGroupInventoryView `json:"inventories,omitempty"`
}

// ChangeIAM2VirtualIDGroupStateEventView ChangeIAM2VirtualIDGroupStateEvent
type ChangeIAM2VirtualIDGroupStateEventView struct {
	Inventory IAM2VirtualIDGroupInventoryView `json:"inventory,omitempty"`
}

// CreateIAM2VirtualIDGroupEventView CreateIAM2VirtualIDGroupEvent
type CreateIAM2VirtualIDGroupEventView struct {
	Inventory IAM2VirtualIDGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteIAM2VirtualIDGroupEventView DeleteIAM2VirtualIDGroupEvent
type DeleteIAM2VirtualIDGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateIAM2VirtualIDGroupEventView UpdateIAM2VirtualIDGroupEvent
type UpdateIAM2VirtualIDGroupEventView struct {
	Inventory IAM2VirtualIDGroupInventoryView `json:"inventory,omitempty"`
}
