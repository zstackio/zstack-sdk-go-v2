// Copyright (c) ZStack.io, Inc.

package param

// DeleteLicenseDetailParam DeleteLicense detail param
type DeleteLicenseDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	Module string `json:"module,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid" validate:"required"`
}

// DeleteLicenseParam DeleteLicense request param
type DeleteLicenseParam struct {
	BaseParam
	Params DeleteLicenseDetailParam `json:"params"`
}
