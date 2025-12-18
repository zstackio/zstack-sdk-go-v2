// Copyright (c) ZStack.io, Inc.

package param

// DeleteBareMetal2GatewayDetailParam DeleteBareMetal2Gateway detail param
type DeleteBareMetal2GatewayDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBareMetal2GatewayParam DeleteBareMetal2Gateway request param
type DeleteBareMetal2GatewayParam struct {
	BaseParam
	Params DeleteBareMetal2GatewayDetailParam `json:"params"`
}
