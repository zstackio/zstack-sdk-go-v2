// Copyright (c) ZStack.io, Inc.

package param

// DetachHybridKeyDetailParam DetachHybridKey detail param
type DetachHybridKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DetachHybridKeyParam DetachHybridKey request param
type DetachHybridKeyParam struct {
	BaseParam
	Params DetachHybridKeyDetailParam `json:"params"`
}
