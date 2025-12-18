// Copyright (c) ZStack.io, Inc.

package param

// DeleteVipDetailParam DeleteVip detail param
type DeleteVipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVipParam DeleteVip request param
type DeleteVipParam struct {
	BaseParam
	Params DeleteVipDetailParam `json:"params"`
}
