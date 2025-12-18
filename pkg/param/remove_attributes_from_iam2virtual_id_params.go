// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2VirtualIDDetailParam RemoveAttributesFromIAM2VirtualID详细参数
type RemoveAttributesFromIAM2VirtualIDDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"attributeUuids" validate:"required"` // 必填
}

// RemoveAttributesFromIAM2VirtualIDParam RemoveAttributesFromIAM2VirtualID请求参数
type RemoveAttributesFromIAM2VirtualIDParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2VirtualIDDetailParam `json:"params"` // 详细参数
}

