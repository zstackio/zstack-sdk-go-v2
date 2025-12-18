// Copyright (c) ZStack.io, Inc.

package param

// UnregisterLicenseRequestedApplicationDetailParam UnregisterLicenseRequestedApplication detail param
type UnregisterLicenseRequestedApplicationDetailParam struct {
	AppId string `json:"appId" validate:"required"`
}

// UnregisterLicenseRequestedApplicationParam UnregisterLicenseRequestedApplication request param
type UnregisterLicenseRequestedApplicationParam struct {
	BaseParam
	Params UnregisterLicenseRequestedApplicationDetailParam `json:"params"`
}
