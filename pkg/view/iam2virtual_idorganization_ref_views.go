// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDOrganizationRefInventoryView IAM2VirtualIDOrganizationRef
type IAM2VirtualIDOrganizationRefInventoryView struct {
	VirtualIDUuid string `json:"virtualIDUuid,omitempty"`
	OrganizationUuid string `json:"organizationUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

