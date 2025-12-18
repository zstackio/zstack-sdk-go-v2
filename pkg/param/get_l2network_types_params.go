// Copyright (c) ZStack.io, Inc.

package param

// GetL2NetworkTypesDetailParam GetL2NetworkTypes detail param
type GetL2NetworkTypesDetailParam struct {
}

// GetL2NetworkTypesParam GetL2NetworkTypes request param
type GetL2NetworkTypesParam struct {
	BaseParam
	Params GetL2NetworkTypesDetailParam `json:"params"`
}
