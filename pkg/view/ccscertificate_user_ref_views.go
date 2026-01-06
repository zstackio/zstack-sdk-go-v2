// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CCSCertificateUserRefInventoryView CCSCertificateUserRef
type CCSCertificateUserRefInventoryView struct {
	UserUuid string `json:"userUuid,omitempty"`
	CertificateUuid string `json:"certificateUuid,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

