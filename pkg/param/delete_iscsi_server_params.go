// Copyright (c) ZStack.io, Inc.

package param

// DeleteIscsiServerDetailParam DeleteIscsiServer detail param
type DeleteIscsiServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIscsiServerParam DeleteIscsiServer request param
type DeleteIscsiServerParam struct {
	BaseParam
	Params DeleteIscsiServerDetailParam `json:"params"`
}
