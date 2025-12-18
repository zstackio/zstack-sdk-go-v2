// Copyright (c) ZStack.io, Inc.

package param

// DeleteVirtualBorderRouterLocalDetailParam DeleteVirtualBorderRouterLocal详细参数
type DeleteVirtualBorderRouterLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVirtualBorderRouterLocalParam DeleteVirtualBorderRouterLocal请求参数
type DeleteVirtualBorderRouterLocalParam struct {
	BaseParam
	Params DeleteVirtualBorderRouterLocalDetailParam `json:"params"` // 详细参数
}

