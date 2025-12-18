// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkTypesDetailParam GetL3NetworkTypes detail param
type GetL3NetworkTypesDetailParam struct {
}

// GetL3NetworkTypesParam GetL3NetworkTypes request param
type GetL3NetworkTypesParam struct {
	BaseParam
	Params GetL3NetworkTypesDetailParam `json:"params"`
}
