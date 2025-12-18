// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CCSCertificateUserRefInventoryView CCSCertificateUserRef
type CCSCertificateUserRefInventoryView struct {
	rest string `json:"userUuid,omitempty"`
	rest string `json:"certificateUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

