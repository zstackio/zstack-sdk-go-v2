// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmNicMacDetailParam UpdateVmNicMac详细参数
type UpdateVmNicMacDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"mac" validate:"required"` // 必填
}

// UpdateVmNicMacParam UpdateVmNicMac请求参数
type UpdateVmNicMacParam struct {
	BaseParam
	Params UpdateVmNicMacDetailParam `json:"params"` // 详细参数
}

