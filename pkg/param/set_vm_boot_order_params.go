// Copyright (c) ZStack.io, Inc.

package param

// SetVmBootOrderDetailParam SetVmBootOrder详细参数
type SetVmBootOrderDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"bootOrder,omitempty"`
}

// SetVmBootOrderParam SetVmBootOrder请求参数
type SetVmBootOrderParam struct {
	BaseParam
	Params SetVmBootOrderDetailParam `json:"params"` // 详细参数
}

