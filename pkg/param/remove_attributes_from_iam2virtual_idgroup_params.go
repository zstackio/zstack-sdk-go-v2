// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2VirtualIDGroupDetailParam RemoveAttributesFromIAM2VirtualIDGroup详细参数
type RemoveAttributesFromIAM2VirtualIDGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"attributeUuids" validate:"required"` // 必填
}

// RemoveAttributesFromIAM2VirtualIDGroupParam RemoveAttributesFromIAM2VirtualIDGroup请求参数
type RemoveAttributesFromIAM2VirtualIDGroupParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2VirtualIDGroupDetailParam `json:"params"` // 详细参数
}

