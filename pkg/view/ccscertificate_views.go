// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CCSCertificateInventoryView CCSCertificate
type CCSCertificateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Algorithm string `json:"algorithm,omitempty"`
	Format string `json:"format,omitempty"`
	IssuerDN string `json:"issuerDN,omitempty"`
	SubjectDN string `json:"subjectDN,omitempty"`
	SerNumber string `json:"serNumber,omitempty"`
	EffectiveTime time.Time `json:"effectiveTime,omitempty"`
	ExpirationTime time.Time `json:"expirationTime,omitempty"`
	UserCertificateRefs []CCSCertificateUserRefInventoryView `json:"userCertificateRefs,omitempty"`
}

// AddCCSCertificateEventView AddCCSCertificateEvent
type AddCCSCertificateEventView struct {
	Inventory CCSCertificateInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteCCSCertificateEventView DeleteCCSCertificateEvent
type DeleteCCSCertificateEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateCCSCertificateUserStateEventView UpdateCCSCertificateUserStateEvent
type UpdateCCSCertificateUserStateEventView struct {
	Inventory CCSCertificateInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AttachCCSCertificateToUserEventView AttachCCSCertificateToUserEvent
type AttachCCSCertificateToUserEventView struct {
	Inventory CCSCertificateInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// QueryCCSCertificateView QueryCCSCertificate
type QueryCCSCertificateView struct {
	Inventories []CCSCertificateInventoryView `json:"inventories,omitempty"`
}

