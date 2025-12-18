// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CCSCertificateInventoryView CCSCertificate
type CCSCertificateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"algorithm,omitempty"`
	rest string `json:"format,omitempty"`
	rest string `json:"issuerDN,omitempty"`
	rest string `json:"subjectDN,omitempty"`
	rest string `json:"serNumber,omitempty"`
	rest time.Time `json:"effectiveTime,omitempty"`
	rest time.Time `json:"expirationTime,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []CCSCertificateUserRefInventoryView `json:"userCertificateRefs,omitempty"`
}

