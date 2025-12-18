// Copyright (c) ZStack.io, Inc.

package param

// DeleteLongJobDetailParam DeleteLongJob detail param
type DeleteLongJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteLongJobParam DeleteLongJob request param
type DeleteLongJobParam struct {
	BaseParam
	Params DeleteLongJobDetailParam `json:"params"`
}
