// Copyright (c) ZStack.io, Inc.

package param

// RefreshIscsiServerDetailParam RefreshIscsiServer detail param
type RefreshIscsiServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RefreshIscsiServerParam RefreshIscsiServer request param
type RefreshIscsiServerParam struct {
	BaseParam
	Params RefreshIscsiServerDetailParam `json:"params"`
}
