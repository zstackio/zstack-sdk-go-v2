// Copyright (c) ZStack.io, Inc.

package param

// GetCreateEcsImageProgressDetailParam GetCreateEcsImageProgress详细参数
type GetCreateEcsImageProgressDetailParam struct {
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
}

// GetCreateEcsImageProgressParam GetCreateEcsImageProgress请求参数
type GetCreateEcsImageProgressParam struct {
	BaseParam
	Params GetCreateEcsImageProgressDetailParam `json:"params"` // 详细参数
}

