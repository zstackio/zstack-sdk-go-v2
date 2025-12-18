// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2VirtualIDOrganizationRefInventoryView IAM2VirtualIDOrganizationRef
type IAM2VirtualIDOrganizationRefInventoryView struct {
	rest string `json:"virtualIDUuid,omitempty"`
	rest string `json:"organizationUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

