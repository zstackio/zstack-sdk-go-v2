// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2OrganizationInventoryView IAM2Organization
type IAM2OrganizationInventoryView struct {
	BaseInfoView
	BaseTimeView
	State                string                       `json:"state,omitempty"`
	Type                 string                       `json:"type,omitempty"`
	SrcType              string                       `json:"srcType,omitempty"`
	ParentUuid           string                       `json:"parentUuid,omitempty"`
	RootOrganizationUuid string                       `json:"rootOrganizationUuid,omitempty"`
	Attributes           []IAM2AttributeInventoryView `json:"attributes,omitempty"`
}

// ChangeIAM2OrganizationStateEventView ChangeIAM2OrganizationStateEvent
type ChangeIAM2OrganizationStateEventView struct {
	Inventory IAM2OrganizationInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2OrganizationView QueryIAM2Organization
type QueryIAM2OrganizationView struct {
	Inventories []IAM2OrganizationInventoryView `json:"inventories,omitempty"`
}

// DeleteIAM2OrganizationEventView DeleteIAM2OrganizationEvent
type DeleteIAM2OrganizationEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateIAM2OrganizationEventView UpdateIAM2OrganizationEvent
type UpdateIAM2OrganizationEventView struct {
	Inventory IAM2OrganizationInventoryView `json:"inventory,omitempty"`
}

// CreateIAM2OrganizationEventView CreateIAM2OrganizationEvent
type CreateIAM2OrganizationEventView struct {
	Inventory IAM2OrganizationInventoryView `json:"inventory,omitempty"`
}
