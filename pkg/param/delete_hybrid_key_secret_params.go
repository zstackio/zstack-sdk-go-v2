// Copyright (c) ZStack.io, Inc.

package param

// DeleteHybridKeySecretDetailParam DeleteHybridKeySecret detail param
type DeleteHybridKeySecretDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHybridKeySecretParam DeleteHybridKeySecret request param
type DeleteHybridKeySecretParam struct {
	BaseParam
	Params DeleteHybridKeySecretDetailParam `json:"params"`
}
