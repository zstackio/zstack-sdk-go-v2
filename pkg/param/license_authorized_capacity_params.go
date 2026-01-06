// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// GetLicenseAuthorizedCapacityParamDetail GetLicenseAuthorizedCapacity detail param
type GetLicenseAuthorizedCapacityParamDetail struct {
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid,omitempty"`
	ShowServerCapacity bool `json:"showServerCapacity,omitempty"`
}

// GetLicenseAuthorizedCapacityParam GetLicenseAuthorizedCapacity request param
type GetLicenseAuthorizedCapacityParam struct {
	BaseParam
	Params GetLicenseAuthorizedCapacityParamDetail `json:"params"`
}
