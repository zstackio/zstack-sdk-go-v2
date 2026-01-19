// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteLicenseParamDetail DeleteLicense detail param
type DeleteLicenseParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	Module *string `json:"module,omitempty"`
}

// DeleteLicenseParam DeleteLicense request param
type DeleteLicenseParam struct {
	BaseParam
	Params DeleteLicenseParamDetail `json:"deleteLicense"`
}
// UpdateLicenseParamDetail UpdateLicense detail param
type UpdateLicenseParamDetail struct {
	License string `json:"license" validate:"required"`
	AdditionSession *string `json:"additionSession,omitempty"`
}

// UpdateLicenseParam UpdateLicense request param
type UpdateLicenseParam struct {
	BaseParam
	Params UpdateLicenseParamDetail `json:"updateLicense"`
}
