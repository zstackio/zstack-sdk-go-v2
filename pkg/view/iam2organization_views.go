// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2OrganizationInventoryView IAM2Organization
type IAM2OrganizationInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	SrcType string `json:"srcType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	RootOrganizationUuid string `json:"rootOrganizationUuid,omitempty"`
	Attributes []IAM2AttributeInventoryView `json:"attributes,omitempty"`
}

