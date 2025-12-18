// Copyright (c) ZStack.io, Inc.

package param

// GetResourceStackVmStatusDetailParam GetResourceStackVmStatus详细参数
type GetResourceStackVmStatusDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetResourceStackVmStatusParam GetResourceStackVmStatus请求参数
type GetResourceStackVmStatusParam struct {
	BaseParam
	Params GetResourceStackVmStatusDetailParam `json:"params"` // 详细参数
}

