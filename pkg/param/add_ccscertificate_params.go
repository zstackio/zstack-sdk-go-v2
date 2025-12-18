// Copyright (c) ZStack.io, Inc.

package param

// AddCCSCertificateDetailParam AddCCSCertificate详细参数
type AddCCSCertificateDetailParam struct {
	rest string `json:"certificate" validate:"required"` // 必填
}

// AddCCSCertificateParam AddCCSCertificate请求参数
type AddCCSCertificateParam struct {
	BaseParam
	Params AddCCSCertificateDetailParam `json:"params"` // 详细参数
}

