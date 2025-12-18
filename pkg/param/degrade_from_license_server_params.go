// Copyright (c) ZStack.io, Inc.

package param

// DegradeFromLicenseServerDetailParam DegradeFromLicenseServer detail param
type DegradeFromLicenseServerDetailParam struct {
}

// DegradeFromLicenseServerParam DegradeFromLicenseServer request param
type DegradeFromLicenseServerParam struct {
	BaseParam
	Params DegradeFromLicenseServerDetailParam `json:"params"`
}
