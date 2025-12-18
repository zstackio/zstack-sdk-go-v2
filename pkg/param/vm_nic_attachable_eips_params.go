// Copyright (c) ZStack.io, Inc.

package param

// GetVmNicAttachableEipsDetailParam GetVmNicAttachableEips详细参数
type GetVmNicAttachableEipsDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest int `json:"ipVersion,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVmNicAttachableEipsParam GetVmNicAttachableEips请求参数
type GetVmNicAttachableEipsParam struct {
	BaseParam
	Params GetVmNicAttachableEipsDetailParam `json:"params"` // 详细参数
}

