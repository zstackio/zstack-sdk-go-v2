// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2OrganizationDetailParam AddAttributesToIAM2Organization detail param
type AddAttributesToIAM2OrganizationDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Attributes []AttributeParam `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2OrganizationParam AddAttributesToIAM2Organization request param
type AddAttributesToIAM2OrganizationParam struct {
	BaseParam
	Params AddAttributesToIAM2OrganizationDetailParam `json:"params"`
}
