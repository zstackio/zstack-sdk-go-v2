// Copyright (c) ZStack.io, Inc.

package param

// DeleteCCSCertificateDetailParam DeleteCCSCertificate详细参数
type DeleteCCSCertificateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteCCSCertificateParam DeleteCCSCertificate请求参数
type DeleteCCSCertificateParam struct {
	BaseParam
	Params DeleteCCSCertificateDetailParam `json:"params"` // 详细参数
}

