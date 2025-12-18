// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CCSCertificateInventoryView CCSCertificate
type CCSCertificateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
	Format string `json:"format,omitempty"`
	IssuerDN string `json:"issuerDN,omitempty"`
	SubjectDN string `json:"subjectDN,omitempty"`
	SerNumber string `json:"serNumber,omitempty"`
	EffectiveTime time.Time `json:"effectiveTime,omitempty"`
	ExpirationTime time.Time `json:"expirationTime,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	UserCertificateRefs []CCSCertificateUserRefInventoryView `json:"userCertificateRefs,omitempty"`
}

