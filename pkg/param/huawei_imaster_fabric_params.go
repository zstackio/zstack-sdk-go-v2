// Copyright (c) ZStack.io, Inc.

package param

// DeleteHuaweiIMasterFabricDetailParam DeleteHuaweiIMasterFabric详细参数
type DeleteHuaweiIMasterFabricDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"sdnControllerUuid,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterFabricParam DeleteHuaweiIMasterFabric请求参数
type DeleteHuaweiIMasterFabricParam struct {
	BaseParam
	Params DeleteHuaweiIMasterFabricDetailParam `json:"params"` // 详细参数
}

