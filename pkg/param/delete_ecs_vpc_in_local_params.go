// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsVpcInLocalDetailParam DeleteEcsVpcInLocal detail param
type DeleteEcsVpcInLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsVpcInLocalParam DeleteEcsVpcInLocal request param
type DeleteEcsVpcInLocalParam struct {
	BaseParam
	Params DeleteEcsVpcInLocalDetailParam `json:"params"`
}
