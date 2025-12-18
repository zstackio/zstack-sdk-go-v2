// Copyright (c) ZStack.io, Inc.

package param

// GetResourceStackFromResourceDetailParam GetResourceStackFromResource详细参数
type GetResourceStackFromResourceDetailParam struct {
	rest string `json:"resourceUuid" validate:"required"` // 必填
}

// GetResourceStackFromResourceParam GetResourceStackFromResource请求参数
type GetResourceStackFromResourceParam struct {
	BaseParam
	Params GetResourceStackFromResourceDetailParam `json:"params"` // 详细参数
}

