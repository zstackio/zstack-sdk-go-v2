// Copyright (c) ZStack.io, Inc.

package param

// DeleteHuaweiIMasterVpcDetailParam DeleteHuaweiIMasterVpc detail param
type DeleteHuaweiIMasterVpcDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterVpcParam DeleteHuaweiIMasterVpc request param
type DeleteHuaweiIMasterVpcParam struct {
	BaseParam
	Params DeleteHuaweiIMasterVpcDetailParam `json:"params"`
}
