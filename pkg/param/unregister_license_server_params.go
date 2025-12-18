// Copyright (c) ZStack.io, Inc.

package param

// UnregisterLicenseServerDetailParam UnregisterLicenseServer detail param
type UnregisterLicenseServerDetailParam struct {
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid,omitempty"`
}

// UnregisterLicenseServerParam UnregisterLicenseServer request param
type UnregisterLicenseServerParam struct {
	BaseParam
	Params UnregisterLicenseServerDetailParam `json:"params"`
}
