// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2OrganizationAttributeInventoryView IAM2OrganizationAttribute
type IAM2OrganizationAttributeInventoryView struct {
	rest string `json:"organizationUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"value,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

