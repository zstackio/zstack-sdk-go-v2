// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBaremetalPxeServerDetailParam ReconnectBaremetalPxeServer detail param
type ReconnectBaremetalPxeServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectBaremetalPxeServerParam ReconnectBaremetalPxeServer request param
type ReconnectBaremetalPxeServerParam struct {
	BaseParam
	Params ReconnectBaremetalPxeServerDetailParam `json:"params"`
}
