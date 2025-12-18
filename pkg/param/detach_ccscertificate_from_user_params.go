// Copyright (c) ZStack.io, Inc.

package param

// DetachCCSCertificateFromUserDetailParam DetachCCSCertificateFromUser detail param
type DetachCCSCertificateFromUserDetailParam struct {
	UserUuid string `json:"userUuid" validate:"required"`
}

// DetachCCSCertificateFromUserParam DetachCCSCertificateFromUser request param
type DetachCCSCertificateFromUserParam struct {
	BaseParam
	Params DetachCCSCertificateFromUserDetailParam `json:"params"`
}
