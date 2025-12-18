// Copyright (c) ZStack.io, Inc.

package param

// DeleteHuaweiIMasterFabricDetailParam DeleteHuaweiIMasterFabric detail param
type DeleteHuaweiIMasterFabricDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterFabricParam DeleteHuaweiIMasterFabric request param
type DeleteHuaweiIMasterFabricParam struct {
	BaseParam
	Params DeleteHuaweiIMasterFabricDetailParam `json:"params"`
}
