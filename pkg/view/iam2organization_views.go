// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2OrganizationInventoryView IAM2Organization
type IAM2OrganizationInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"srcType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"parentUuid,omitempty"`
	rest string `json:"rootOrganizationUuid,omitempty"`
	rest []IAM2AttributeInventoryView `json:"attributes,omitempty"`
}

