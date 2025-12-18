// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2OrganizationProjectRefInventoryView IAM2OrganizationProjectRef
type IAM2OrganizationProjectRefInventoryView struct {
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"organizationUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

