// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteHuaweiIMasterFabricParamDetail DeleteHuaweiIMasterFabric detail param
type DeleteHuaweiIMasterFabricParamDetail struct {
	SdnControllerUuid *string `json:"sdnControllerUuid,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterFabricParam DeleteHuaweiIMasterFabric request param
type DeleteHuaweiIMasterFabricParam struct {
	BaseParam
	Params DeleteHuaweiIMasterFabricParamDetail `json:"deleteHuaweiIMasterFabric"`
}
