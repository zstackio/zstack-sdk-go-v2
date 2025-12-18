// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunRouterInterfaceRemoteDetailParam DeleteAliyunRouterInterfaceRemote detail param
type DeleteAliyunRouterInterfaceRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VRouterType string `json:"vRouterType" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouterInterfaceRemoteParam DeleteAliyunRouterInterfaceRemote request param
type DeleteAliyunRouterInterfaceRemoteParam struct {
	BaseParam
	Params DeleteAliyunRouterInterfaceRemoteDetailParam `json:"params"`
}
