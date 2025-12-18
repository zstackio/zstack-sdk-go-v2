// Copyright (c) ZStack.io, Inc.

package param

// GetVmBootOrderDetailParam GetVmBootOrder详细参数
type GetVmBootOrderDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmBootOrderParam GetVmBootOrder请求参数
type GetVmBootOrderParam struct {
	BaseParam
	Params GetVmBootOrderDetailParam `json:"params"` // 详细参数
}

