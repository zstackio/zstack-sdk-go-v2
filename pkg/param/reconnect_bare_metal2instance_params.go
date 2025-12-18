// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBareMetal2InstanceDetailParam ReconnectBareMetal2Instance detail param
type ReconnectBareMetal2InstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectBareMetal2InstanceParam ReconnectBareMetal2Instance request param
type ReconnectBareMetal2InstanceParam struct {
	BaseParam
	Params ReconnectBareMetal2InstanceDetailParam `json:"params"`
}
