// Copyright (c) ZStack.io, Inc.

package param

// UpdateMdevDeviceDetailParam UpdateMdevDevice详细参数
type UpdateMdevDeviceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
}

// UpdateMdevDeviceParam UpdateMdevDevice请求参数
type UpdateMdevDeviceParam struct {
	BaseParam
	Params UpdateMdevDeviceDetailParam `json:"params"` // 详细参数
}

