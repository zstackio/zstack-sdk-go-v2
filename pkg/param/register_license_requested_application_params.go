// Copyright (c) ZStack.io, Inc.

package param

// RegisterLicenseRequestedApplicationDetailParam RegisterLicenseRequestedApplication detail param
type RegisterLicenseRequestedApplicationDetailParam struct {
	LicenseRequestCode string `json:"licenseRequestCode" validate:"required"`
	ClientPubKey string `json:"clientPubKey,omitempty"`
	CurrentTimeMillis int64 `json:"currentTimeMillis,omitempty"`
}

// RegisterLicenseRequestedApplicationParam RegisterLicenseRequestedApplication request param
type RegisterLicenseRequestedApplicationParam struct {
	BaseParam
	Params RegisterLicenseRequestedApplicationDetailParam `json:"params"`
}
