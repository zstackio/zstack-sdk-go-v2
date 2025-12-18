// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunRouterInterfaceLocalDetailParam DeleteAliyunRouterInterfaceLocal详细参数
type DeleteAliyunRouterInterfaceLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouterInterfaceLocalParam DeleteAliyunRouterInterfaceLocal请求参数
type DeleteAliyunRouterInterfaceLocalParam struct {
	BaseParam
	Params DeleteAliyunRouterInterfaceLocalDetailParam `json:"params"` // 详细参数
}

