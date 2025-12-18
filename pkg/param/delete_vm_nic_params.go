// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmNicDetailParam DeleteVmNic detail param
type DeleteVmNicDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmNicParam DeleteVmNic request param
type DeleteVmNicParam struct {
	BaseParam
	Params DeleteVmNicDetailParam `json:"params"`
}
