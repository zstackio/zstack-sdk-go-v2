// Copyright (c) ZStack.io, Inc.

package view

// UpdateCCSCertificateUserStateEventView UpdateCCSCertificateUserStateEvent
type UpdateCCSCertificateUserStateEventView struct {
	Inventory CCSCertificateInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

