// Copyright (c) ZStack.io, Inc.

package param

// AttachL3NetworkToVmNicDetailParam AttachL3NetworkToVmNic详细参数
type AttachL3NetworkToVmNicDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"staticIp,omitempty"`
}

// AttachL3NetworkToVmNicParam AttachL3NetworkToVmNic请求参数
type AttachL3NetworkToVmNicParam struct {
	BaseParam
	Params AttachL3NetworkToVmNicDetailParam `json:"params"` // 详细参数
}

