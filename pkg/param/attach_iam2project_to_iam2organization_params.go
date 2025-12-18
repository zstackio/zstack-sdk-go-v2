// Copyright (c) ZStack.io, Inc.

package param

// AttachIAM2ProjectToIAM2OrganizationDetailParam AttachIAM2ProjectToIAM2Organization detail param
type AttachIAM2ProjectToIAM2OrganizationDetailParam struct {
	ProjectUuid string `json:"projectUuid" validate:"required"`
	OrganizationUuid string `json:"organizationUuid" validate:"required"`
}

// AttachIAM2ProjectToIAM2OrganizationParam AttachIAM2ProjectToIAM2Organization request param
type AttachIAM2ProjectToIAM2OrganizationParam struct {
	BaseParam
	Params AttachIAM2ProjectToIAM2OrganizationDetailParam `json:"params"`
}
