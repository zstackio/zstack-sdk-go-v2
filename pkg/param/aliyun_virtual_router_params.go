// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunVirtualRouterDetailParam UpdateAliyunVirtualRouter详细参数
type UpdateAliyunVirtualRouterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateAliyunVirtualRouterParam UpdateAliyunVirtualRouter请求参数
type UpdateAliyunVirtualRouterParam struct {
	BaseParam
	Params UpdateAliyunVirtualRouterDetailParam `json:"params"` // 详细参数
}

