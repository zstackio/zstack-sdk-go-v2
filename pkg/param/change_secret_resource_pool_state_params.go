// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecretResourcePoolStateDetailParam ChangeSecretResourcePoolState detail param
type ChangeSecretResourcePoolStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSecretResourcePoolStateParam ChangeSecretResourcePoolState request param
type ChangeSecretResourcePoolStateParam struct {
	BaseParam
	Params ChangeSecretResourcePoolStateDetailParam `json:"params"`
}
