// Copyright (c) ZStack.io, Inc.

package param

// UpdateCCSCertificateUserStateDetailParam UpdateCCSCertificateUserState detail param
type UpdateCCSCertificateUserStateDetailParam struct {
	UserUuid string `json:"userUuid" validate:"required"`
	State string `json:"state" validate:"required"`
}

// UpdateCCSCertificateUserStateParam UpdateCCSCertificateUserState request param
type UpdateCCSCertificateUserStateParam struct {
	BaseParam
	Params UpdateCCSCertificateUserStateDetailParam `json:"params"`
}
