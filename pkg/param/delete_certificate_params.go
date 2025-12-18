// Copyright (c) ZStack.io, Inc.

package param

// DeleteCertificateDetailParam DeleteCertificate detail param
type DeleteCertificateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteCertificateParam DeleteCertificate request param
type DeleteCertificateParam struct {
	BaseParam
	Params DeleteCertificateDetailParam `json:"params"`
}
