// Copyright (c) ZStack.io, Inc.

package param

// GetVirtualizerInfoDetailParam GetVirtualizerInfo详细参数
type GetVirtualizerInfoDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
}

// GetVirtualizerInfoParam GetVirtualizerInfo请求参数
type GetVirtualizerInfoParam struct {
	BaseParam
	Params GetVirtualizerInfoDetailParam `json:"params"` // 详细参数
}

