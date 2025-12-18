// Copyright (c) ZStack.io, Inc.

package param

// GetHostIommuStateDetailParam GetHostIommuState详细参数
type GetHostIommuStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetHostIommuStateParam GetHostIommuState请求参数
type GetHostIommuStateParam struct {
	BaseParam
	Params GetHostIommuStateDetailParam `json:"params"` // 详细参数
}

