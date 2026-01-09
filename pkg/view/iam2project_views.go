// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectInventoryView IAM2Project
type IAM2ProjectInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	Attributes []IAM2AttributeInventoryView `json:"attributes,omitempty"`
	LinkedAccountUuid *string `json:"linkedAccountUuid,omitempty"`
}

// CreateIAM2ProjectEventView CreateIAM2ProjectEvent
type CreateIAM2ProjectEventView struct {
	Inventory IAM2ProjectInventoryView `json:"inventory,omitempty"`
}

// DetachIAM2ProjectFromIAM2OrganizationEventView DetachIAM2ProjectFromIAM2OrganizationEvent
type DetachIAM2ProjectFromIAM2OrganizationEventView struct {
	Inventory IAM2ProjectInventoryView `json:"inventory,omitempty"`
}

// DeleteIAM2ProjectEventView DeleteIAM2ProjectEvent
type DeleteIAM2ProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateIAM2ProjectFromTemplateEventView CreateIAM2ProjectFromTemplateEvent
type CreateIAM2ProjectFromTemplateEventView struct {
	Inventory IAM2ProjectInventoryView `json:"inventory,omitempty"`
}

// RecoverIAM2ProjectEventView RecoverIAM2ProjectEvent
type RecoverIAM2ProjectEventView struct {
	Inventory IAM2ProjectInventoryView `json:"inventory,omitempty"`
}

// AttachIAM2ProjectToIAM2OrganizationEventView AttachIAM2ProjectToIAM2OrganizationEvent
type AttachIAM2ProjectToIAM2OrganizationEventView struct {
	Inventory IAM2ProjectInventoryView `json:"inventory,omitempty"`
}

// ChangeIAM2ProjectStateEventView ChangeIAM2ProjectStateEvent
type ChangeIAM2ProjectStateEventView struct {
	Inventory IAM2ProjectInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2ProjectView QueryIAM2Project
type QueryIAM2ProjectView struct {
	Inventories []IAM2ProjectInventoryView `json:"inventories,omitempty"`
}

// GetIAM2ProjectsOfVirtualIDView GetIAM2ProjectsOfVirtualID
type GetIAM2ProjectsOfVirtualIDView struct {
	Inventories []IAM2ProjectInventoryView `json:"inventories,omitempty"`
}

// ExpungeIAM2ProjectEventView ExpungeIAM2ProjectEvent
type ExpungeIAM2ProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// LoginIAM2ProjectView LoginIAM2Project
type LoginIAM2ProjectView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// UpdateIAM2ProjectEventView UpdateIAM2ProjectEvent
type UpdateIAM2ProjectEventView struct {
	Inventory IAM2ProjectInventoryView `json:"inventory,omitempty"`
}

