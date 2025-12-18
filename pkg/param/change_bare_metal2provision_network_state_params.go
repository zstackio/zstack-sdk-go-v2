// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2ProvisionNetworkStateDetailParam ChangeBareMetal2ProvisionNetworkState detail param
type ChangeBareMetal2ProvisionNetworkStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2ProvisionNetworkStateParam ChangeBareMetal2ProvisionNetworkState request param
type ChangeBareMetal2ProvisionNetworkStateParam struct {
	BaseParam
	Params ChangeBareMetal2ProvisionNetworkStateDetailParam `json:"params"`
}
