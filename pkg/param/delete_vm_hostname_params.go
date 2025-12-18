// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmHostnameDetailParam DeleteVmHostname detail param
type DeleteVmHostnameDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmHostnameParam DeleteVmHostname request param
type DeleteVmHostnameParam struct {
	BaseParam
	Params DeleteVmHostnameDetailParam `json:"params"`
}
