// Copyright (c) ZStack.io, Inc.

package param

// IdentifyHostDetailParam IdentifyHost detail param
type IdentifyHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Interval int64 `json:"interval,omitempty"`
}

// IdentifyHostParam IdentifyHost request param
type IdentifyHostParam struct {
	BaseParam
	Params IdentifyHostDetailParam `json:"params"`
}
