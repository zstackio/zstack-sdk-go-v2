// Copyright (c) ZStack.io, Inc.

package param

// DeleteSecretResourcePoolDetailParam DeleteSecretResourcePool detail param
type DeleteSecretResourcePoolDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSecretResourcePoolParam DeleteSecretResourcePool request param
type DeleteSecretResourcePoolParam struct {
	BaseParam
	Params DeleteSecretResourcePoolDetailParam `json:"params"`
}
