// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IAM2OrganizationAttributeInventoryView IAM2OrganizationAttribute
type IAM2OrganizationAttributeInventoryView struct {
	BaseInfoView
	BaseTimeView
	OrganizationUuid string `json:"organizationUuid,omitempty"`
	Value string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateIAM2OrganizationAttributeEventView UpdateIAM2OrganizationAttributeEvent
type UpdateIAM2OrganizationAttributeEventView struct {
	Inventory IAM2OrganizationAttributeInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2OrganizationAttributeView QueryIAM2OrganizationAttribute
type QueryIAM2OrganizationAttributeView struct {
	Inventories []IAM2OrganizationAttributeInventoryView `json:"inventories,omitempty"`
}

