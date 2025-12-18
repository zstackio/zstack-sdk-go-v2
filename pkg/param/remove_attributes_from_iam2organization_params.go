// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2OrganizationDetailParam RemoveAttributesFromIAM2Organization详细参数
type RemoveAttributesFromIAM2OrganizationDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"attributeUuids" validate:"required"` // 必填
}

// RemoveAttributesFromIAM2OrganizationParam RemoveAttributesFromIAM2Organization请求参数
type RemoveAttributesFromIAM2OrganizationParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2OrganizationDetailParam `json:"params"` // 详细参数
}

