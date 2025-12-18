// Copyright (c) ZStack.io, Inc.

package param

// DeleteBareMetal2ProvisionNetworkDetailParam DeleteBareMetal2ProvisionNetwork detail param
type DeleteBareMetal2ProvisionNetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBareMetal2ProvisionNetworkParam DeleteBareMetal2ProvisionNetwork request param
type DeleteBareMetal2ProvisionNetworkParam struct {
	BaseParam
	Params DeleteBareMetal2ProvisionNetworkDetailParam `json:"params"`
}
