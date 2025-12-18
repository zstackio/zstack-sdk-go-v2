// Copyright (c) ZStack.io, Inc.

package param

// DetachIAM2ProjectFromIAM2OrganizationDetailParam DetachIAM2ProjectFromIAM2Organization detail param
type DetachIAM2ProjectFromIAM2OrganizationDetailParam struct {
	ProjectUuid string `json:"projectUuid" validate:"required"`
}

// DetachIAM2ProjectFromIAM2OrganizationParam DetachIAM2ProjectFromIAM2Organization request param
type DetachIAM2ProjectFromIAM2OrganizationParam struct {
	BaseParam
	Params DetachIAM2ProjectFromIAM2OrganizationDetailParam `json:"params"`
}
