// Copyright (c) ZStack.io, Inc.

package param

// AddZBoxDetailParam AddZBox详细参数
type AddZBoxDetailParam struct {
	rest string `json:"usbDeviceUuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest bool `json:"skipFormat,omitempty"`
}

// AddZBoxParam AddZBox请求参数
type AddZBoxParam struct {
	BaseParam
	Params AddZBoxDetailParam `json:"params"` // 详细参数
}

