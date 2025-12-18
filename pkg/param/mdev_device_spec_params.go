// Copyright (c) ZStack.io, Inc.

package param

// UpdateMdevDeviceSpecDetailParam UpdateMdevDeviceSpec详细参数
type UpdateMdevDeviceSpecDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
}

// UpdateMdevDeviceSpecParam UpdateMdevDeviceSpec请求参数
type UpdateMdevDeviceSpecParam struct {
	BaseParam
	Params UpdateMdevDeviceSpecDetailParam `json:"params"` // 详细参数
}

