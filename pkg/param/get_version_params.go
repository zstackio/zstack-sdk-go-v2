// Copyright (c) ZStack.io, Inc.

package param

// GetVersionDetailParam GetVersion detail param
type GetVersionDetailParam struct {
}

// GetVersionParam GetVersion request param
type GetVersionParam struct {
	BaseParam
	Params GetVersionDetailParam `json:"params"`
}
