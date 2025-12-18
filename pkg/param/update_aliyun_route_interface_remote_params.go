// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunRouteInterfaceRemoteDetailParam UpdateAliyunRouteInterfaceRemote detail param
type UpdateAliyunRouteInterfaceRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Op string `json:"op" validate:"required"`
	VRouterType string `json:"vRouterType" validate:"required"`
}

// UpdateAliyunRouteInterfaceRemoteParam UpdateAliyunRouteInterfaceRemote request param
type UpdateAliyunRouteInterfaceRemoteParam struct {
	BaseParam
	Params UpdateAliyunRouteInterfaceRemoteDetailParam `json:"params"`
}
