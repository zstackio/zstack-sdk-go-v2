// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteHuaweiIMasterFabricParamDetail DeleteHuaweiIMasterFabric detail param
type DeleteHuaweiIMasterFabricParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterFabricParam DeleteHuaweiIMasterFabric request param
type DeleteHuaweiIMasterFabricParam struct {
	BaseParam
	DeleteHuaweiIMasterFabric DeleteHuaweiIMasterFabricParamDetail `json:"deleteHuaweiIMasterFabric"`
}
