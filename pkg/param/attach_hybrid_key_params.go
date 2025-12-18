// Copyright (c) ZStack.io, Inc.

package param

// AttachHybridKeyDetailParam AttachHybridKey detail param
type AttachHybridKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// AttachHybridKeyParam AttachHybridKey request param
type AttachHybridKeyParam struct {
	BaseParam
	Params AttachHybridKeyDetailParam `json:"params"`
}
