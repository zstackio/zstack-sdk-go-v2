// Copyright (c) ZStack.io, Inc.

package param

// AddCCSCertificateDetailParam AddCCSCertificate detail param
type AddCCSCertificateDetailParam struct {
	Certificate string `json:"certificate" validate:"required"`
}

// AddCCSCertificateParam AddCCSCertificate request param
type AddCCSCertificateParam struct {
	BaseParam
	Params AddCCSCertificateDetailParam `json:"params"`
}
