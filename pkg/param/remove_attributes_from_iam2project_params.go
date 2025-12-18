// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2ProjectDetailParam RemoveAttributesFromIAM2Project详细参数
type RemoveAttributesFromIAM2ProjectDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"attributeUuids" validate:"required"` // 必填
}

// RemoveAttributesFromIAM2ProjectParam RemoveAttributesFromIAM2Project请求参数
type RemoveAttributesFromIAM2ProjectParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2ProjectDetailParam `json:"params"` // 详细参数
}

