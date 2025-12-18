// Copyright (c) ZStack.io, Inc.

package param

// DeleteIPsecConnectionDetailParam DeleteIPsecConnection detail param
type DeleteIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIPsecConnectionParam DeleteIPsecConnection request param
type DeleteIPsecConnectionParam struct {
	BaseParam
	Params DeleteIPsecConnectionDetailParam `json:"params"`
}
