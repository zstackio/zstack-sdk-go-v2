// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2OrganizationProjectRefInventoryView IAM2OrganizationProjectRef
type IAM2OrganizationProjectRefInventoryView struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	OrganizationUuid string `json:"organizationUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

