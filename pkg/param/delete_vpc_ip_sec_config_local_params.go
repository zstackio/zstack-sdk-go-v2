// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcIpSecConfigLocalDetailParam DeleteVpcIpSecConfigLocal detail param
type DeleteVpcIpSecConfigLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcIpSecConfigLocalParam DeleteVpcIpSecConfigLocal request param
type DeleteVpcIpSecConfigLocalParam struct {
	BaseParam
	Params DeleteVpcIpSecConfigLocalDetailParam `json:"params"`
}
