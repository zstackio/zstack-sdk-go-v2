// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2OrganizationVirtualIDNumberDetailParam GetIAM2OrganizationVirtualIDNumber detail param
type GetIAM2OrganizationVirtualIDNumberDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetIAM2OrganizationVirtualIDNumberParam GetIAM2OrganizationVirtualIDNumber request param
type GetIAM2OrganizationVirtualIDNumberParam struct {
	BaseParam
	Params GetIAM2OrganizationVirtualIDNumberDetailParam `json:"params"`
}
