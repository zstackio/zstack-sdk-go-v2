// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2OrganizationAttributeInventoryView IAM2OrganizationAttribute
type IAM2OrganizationAttributeInventoryView struct {
	OrganizationUuid string `json:"organizationUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// UpdateIAM2OrganizationAttributeEventView UpdateIAM2OrganizationAttributeEvent
type UpdateIAM2OrganizationAttributeEventView struct {
	Inventory IAM2OrganizationAttributeInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2OrganizationAttributeView QueryIAM2OrganizationAttribute
type QueryIAM2OrganizationAttributeView struct {
	Inventories []IAM2OrganizationAttributeInventoryView `json:"inventories,omitempty"`
}

