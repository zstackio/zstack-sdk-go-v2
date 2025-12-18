// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunRouteInterfaceRemoteDetailParam UpdateAliyunRouteInterfaceRemote详细参数
type UpdateAliyunRouteInterfaceRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"op" validate:"required"` // 必填
	rest string `json:"vRouterType" validate:"required"` // 必填
}

// UpdateAliyunRouteInterfaceRemoteParam UpdateAliyunRouteInterfaceRemote请求参数
type UpdateAliyunRouteInterfaceRemoteParam struct {
	BaseParam
	Params UpdateAliyunRouteInterfaceRemoteDetailParam `json:"params"` // 详细参数
}

