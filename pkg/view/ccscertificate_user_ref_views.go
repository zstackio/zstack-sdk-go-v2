// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CCSCertificateUserRefInventoryView CCSCertificateUserRef
type CCSCertificateUserRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	UserUuid string `json:"userUuid,omitempty"`
	CertificateUuid string `json:"certificateUuid,omitempty"`
	State string `json:"state,omitempty"`
}

