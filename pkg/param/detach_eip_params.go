// Copyright (c) ZStack.io, Inc.

package param

// DetachEipDetailParam DetachEip detail param
type DetachEipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DetachEipParam DetachEip request param
type DetachEipParam struct {
	BaseParam
	Params DetachEipDetailParam `json:"params"`
}
