// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunRouterInterfaceLocalDetailParam DeleteAliyunRouterInterfaceLocal detail param
type DeleteAliyunRouterInterfaceLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouterInterfaceLocalParam DeleteAliyunRouterInterfaceLocal request param
type DeleteAliyunRouterInterfaceLocalParam struct {
	BaseParam
	Params DeleteAliyunRouterInterfaceLocalDetailParam `json:"params"`
}
