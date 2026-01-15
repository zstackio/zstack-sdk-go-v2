// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteHuaweiIMasterVpcParamDetail DeleteHuaweiIMasterVpc detail param
type DeleteHuaweiIMasterVpcParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterVpcParam DeleteHuaweiIMasterVpc request param
type DeleteHuaweiIMasterVpcParam struct {
	BaseParam
	Params DeleteHuaweiIMasterVpcParamDetail `json:"deleteHuaweiIMasterVpc"`
}
