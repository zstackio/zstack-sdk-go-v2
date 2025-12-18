// Copyright (c) ZStack.io, Inc.

package param

// UpdateCCSCertificateUserStateDetailParam UpdateCCSCertificateUserState详细参数
type UpdateCCSCertificateUserStateDetailParam struct {
	rest string `json:"userUuid" validate:"required"` // 必填
	rest string `json:"state" validate:"required"` // 必填
}

// UpdateCCSCertificateUserStateParam UpdateCCSCertificateUserState请求参数
type UpdateCCSCertificateUserStateParam struct {
	BaseParam
	Params UpdateCCSCertificateUserStateDetailParam `json:"params"` // 详细参数
}

