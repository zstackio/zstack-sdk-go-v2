// Copyright (c) ZStack.io, Inc.

package param

// UpdateCertificateDetailParam UpdateCertificate详细参数
type UpdateCertificateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// UpdateCertificateParam UpdateCertificate请求参数
type UpdateCertificateParam struct {
	BaseParam
	Params UpdateCertificateDetailParam `json:"params"` // 详细参数
}

