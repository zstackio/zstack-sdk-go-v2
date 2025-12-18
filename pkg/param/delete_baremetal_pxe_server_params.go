// Copyright (c) ZStack.io, Inc.

package param

// DeleteBaremetalPxeServerDetailParam DeleteBaremetalPxeServer detail param
type DeleteBaremetalPxeServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBaremetalPxeServerParam DeleteBaremetalPxeServer request param
type DeleteBaremetalPxeServerParam struct {
	BaseParam
	Params DeleteBaremetalPxeServerDetailParam `json:"params"`
}
