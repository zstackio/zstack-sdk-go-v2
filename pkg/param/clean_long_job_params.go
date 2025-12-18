// Copyright (c) ZStack.io, Inc.

package param

// CleanLongJobDetailParam CleanLongJob detail param
type CleanLongJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// CleanLongJobParam CleanLongJob request param
type CleanLongJobParam struct {
	BaseParam
	Params CleanLongJobDetailParam `json:"params"`
}
