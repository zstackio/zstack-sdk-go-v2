// Copyright (c) ZStack.io, Inc.

package param

// GetResourceAccountDetailParam GetResourceAccount详细参数
type GetResourceAccountDetailParam struct {
	rest []string `json:"resourceUuids" validate:"required"` // 必填
}

// GetResourceAccountParam GetResourceAccount请求参数
type GetResourceAccountParam struct {
	BaseParam
	Params GetResourceAccountDetailParam `json:"params"` // 详细参数
}

