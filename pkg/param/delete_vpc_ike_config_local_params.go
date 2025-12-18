// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcIkeConfigLocalDetailParam DeleteVpcIkeConfigLocal detail param
type DeleteVpcIkeConfigLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcIkeConfigLocalParam DeleteVpcIkeConfigLocal request param
type DeleteVpcIkeConfigLocalParam struct {
	BaseParam
	Params DeleteVpcIkeConfigLocalDetailParam `json:"params"`
}
