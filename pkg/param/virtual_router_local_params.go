// Copyright (c) ZStack.io, Inc.

package param

// DeleteVirtualRouterLocalDetailParam DeleteVirtualRouterLocal详细参数
type DeleteVirtualRouterLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVirtualRouterLocalParam DeleteVirtualRouterLocal请求参数
type DeleteVirtualRouterLocalParam struct {
	BaseParam
	Params DeleteVirtualRouterLocalDetailParam `json:"params"` // 详细参数
}

