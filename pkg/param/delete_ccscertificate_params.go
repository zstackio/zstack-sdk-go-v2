// Copyright (c) ZStack.io, Inc.

package param

// DeleteCCSCertificateDetailParam DeleteCCSCertificate detail param
type DeleteCCSCertificateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteCCSCertificateParam DeleteCCSCertificate request param
type DeleteCCSCertificateParam struct {
	BaseParam
	Params DeleteCCSCertificateDetailParam `json:"params"`
}
