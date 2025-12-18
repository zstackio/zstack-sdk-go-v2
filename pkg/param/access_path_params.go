// Copyright (c) ZStack.io, Inc.

package param

// GetAccessPathDetailParam GetAccessPath详细参数
type GetAccessPathDetailParam struct {
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
}

// GetAccessPathParam GetAccessPath请求参数
type GetAccessPathParam struct {
	BaseParam
	Params GetAccessPathDetailParam `json:"params"` // 详细参数
}

