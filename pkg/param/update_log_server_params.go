// Copyright (c) ZStack.io, Inc.

package param

// UpdateLogServerDetailParam UpdateLogServer detail param
type UpdateLogServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateLogServerParam UpdateLogServer request param
type UpdateLogServerParam struct {
	BaseParam
	Params UpdateLogServerDetailParam `json:"params"`
}
