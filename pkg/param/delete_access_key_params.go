// Copyright (c) ZStack.io, Inc.

package param

// DeleteAccessKeyDetailParam DeleteAccessKey detail param
type DeleteAccessKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAccessKeyParam DeleteAccessKey request param
type DeleteAccessKeyParam struct {
	BaseParam
	Params DeleteAccessKeyDetailParam `json:"params"`
}
