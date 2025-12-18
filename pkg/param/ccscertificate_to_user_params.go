// Copyright (c) ZStack.io, Inc.

package param

// AttachCCSCertificateToUserDetailParam AttachCCSCertificateToUser详细参数
type AttachCCSCertificateToUserDetailParam struct {
	rest string `json:"certificateUuid,omitempty"`
	rest string `json:"userUuid" validate:"required"` // 必填
	rest string `json:"state,omitempty"`
}

// AttachCCSCertificateToUserParam AttachCCSCertificateToUser请求参数
type AttachCCSCertificateToUserParam struct {
	BaseParam
	Params AttachCCSCertificateToUserDetailParam `json:"params"` // 详细参数
}

