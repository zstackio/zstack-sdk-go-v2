// Copyright (c) ZStack.io, Inc.

package param

// CancelLongJobDetailParam CancelLongJob detail param
type CancelLongJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// CancelLongJobParam CancelLongJob request param
type CancelLongJobParam struct {
	BaseParam
	Params CancelLongJobDetailParam `json:"params"`
}
