// Copyright (c) ZStack.io, Inc.

package param

// AddSharedBlockToSharedBlockGroupDetailParam AddSharedBlockToSharedBlockGroup详细参数
type AddSharedBlockToSharedBlockGroupDetailParam struct {
	rest string `json:"diskUuid" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
}

// AddSharedBlockToSharedBlockGroupParam AddSharedBlockToSharedBlockGroup请求参数
type AddSharedBlockToSharedBlockGroupParam struct {
	BaseParam
	Params AddSharedBlockToSharedBlockGroupDetailParam `json:"params"` // 详细参数
}

