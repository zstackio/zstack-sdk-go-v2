// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBareMetal2GatewayDetailParam ReconnectBareMetal2Gateway detail param
type ReconnectBareMetal2GatewayDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectBareMetal2GatewayParam ReconnectBareMetal2Gateway request param
type ReconnectBareMetal2GatewayParam struct {
	BaseParam
	Params ReconnectBareMetal2GatewayDetailParam `json:"params"`
}
