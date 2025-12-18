// Copyright (c) ZStack.io, Inc.

package param

// SetImageQgaDetailParam SetImageQga detail param
type SetImageQgaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Enable bool `json:"enable" validate:"required"`
}

// SetImageQgaParam SetImageQga request param
type SetImageQgaParam struct {
	BaseParam
	Params SetImageQgaDetailParam `json:"params"`
}
