// Copyright (c) ZStack.io, Inc.

package param

// GetResourceFromResourceStackDetailParam GetResourceFromResourceStack详细参数
type GetResourceFromResourceStackDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetResourceFromResourceStackParam GetResourceFromResourceStack请求参数
type GetResourceFromResourceStackParam struct {
	BaseParam
	Params GetResourceFromResourceStackDetailParam `json:"params"` // 详细参数
}

