// Copyright (c) ZStack.io, Inc.

package param

// DeleteHuaweiIMasterVRouterDetailParam DeleteHuaweiIMasterVRouter detail param
type DeleteHuaweiIMasterVRouterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterVRouterParam DeleteHuaweiIMasterVRouter request param
type DeleteHuaweiIMasterVRouterParam struct {
	BaseParam
	Params DeleteHuaweiIMasterVRouterDetailParam `json:"params"`
}
