// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2OrganizationDetailParam DeleteIAM2Organization detail param
type DeleteIAM2OrganizationDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteIAM2OrganizationParam DeleteIAM2Organization request param
type DeleteIAM2OrganizationParam struct {
	BaseParam
	Params DeleteIAM2OrganizationDetailParam `json:"params"`
}
