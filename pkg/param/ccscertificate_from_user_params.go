// Copyright (c) ZStack.io, Inc.

package param

// DetachCCSCertificateFromUserDetailParam DetachCCSCertificateFromUser详细参数
type DetachCCSCertificateFromUserDetailParam struct {
	rest string `json:"userUuid" validate:"required"` // 必填
}

// DetachCCSCertificateFromUserParam DetachCCSCertificateFromUser请求参数
type DetachCCSCertificateFromUserParam struct {
	BaseParam
	Params DetachCCSCertificateFromUserDetailParam `json:"params"` // 详细参数
}

