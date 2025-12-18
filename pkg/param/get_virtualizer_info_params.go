// Copyright (c) ZStack.io, Inc.

package param

// GetVirtualizerInfoDetailParam GetVirtualizerInfo detail param
type GetVirtualizerInfoDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// GetVirtualizerInfoParam GetVirtualizerInfo request param
type GetVirtualizerInfoParam struct {
	BaseParam
	Params GetVirtualizerInfoDetailParam `json:"params"`
}
