// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2OrganizationVirtualIDNumberDetailParam GetIAM2OrganizationVirtualIDNumber详细参数
type GetIAM2OrganizationVirtualIDNumberDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetIAM2OrganizationVirtualIDNumberParam GetIAM2OrganizationVirtualIDNumber请求参数
type GetIAM2OrganizationVirtualIDNumberParam struct {
	BaseParam
	Params GetIAM2OrganizationVirtualIDNumberDetailParam `json:"params"` // 详细参数
}

