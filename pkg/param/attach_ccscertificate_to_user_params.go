// Copyright (c) ZStack.io, Inc.

package param

// AttachCCSCertificateToUserDetailParam AttachCCSCertificateToUser detail param
type AttachCCSCertificateToUserDetailParam struct {
	CertificateUuid string `json:"certificateUuid,omitempty"`
	UserUuid string `json:"userUuid" validate:"required"`
	State string `json:"state,omitempty"`
}

// AttachCCSCertificateToUserParam AttachCCSCertificateToUser request param
type AttachCCSCertificateToUserParam struct {
	BaseParam
	Params AttachCCSCertificateToUserDetailParam `json:"params"`
}
