// Copyright (c) ZStack.io, Inc.

package param

// RegisterLicenseServerDetailParam RegisterLicenseServer detail param
type RegisterLicenseServerDetailParam struct {
	Ip string `json:"ip" validate:"required"`
	LoginParams map[string]interface{} `json:"loginParams" validate:"required"`
}

// RegisterLicenseServerParam RegisterLicenseServer request param
type RegisterLicenseServerParam struct {
	BaseParam
	Params RegisterLicenseServerDetailParam `json:"params"`
}
