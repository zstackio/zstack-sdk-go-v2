// Copyright (c) ZStack.io, Inc.

package param

// ReconnectVirtualRouterDetailParam ReconnectVirtualRouter详细参数
type ReconnectVirtualRouterDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// ReconnectVirtualRouterParam ReconnectVirtualRouter请求参数
type ReconnectVirtualRouterParam struct {
	BaseParam
	Params ReconnectVirtualRouterDetailParam `json:"params"` // 详细参数
}

