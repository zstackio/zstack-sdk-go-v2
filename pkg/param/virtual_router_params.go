// Copyright (c) ZStack.io, Inc.

package param

// UpdateVirtualRouterDetailParam UpdateVirtualRouter详细参数
type UpdateVirtualRouterDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"defaultRouteL3NetworkUuid,omitempty"`
}

// UpdateVirtualRouterParam UpdateVirtualRouter请求参数
type UpdateVirtualRouterParam struct {
	BaseParam
	Params UpdateVirtualRouterDetailParam `json:"params"` // 详细参数
}

