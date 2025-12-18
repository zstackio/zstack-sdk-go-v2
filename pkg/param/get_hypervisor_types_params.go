// Copyright (c) ZStack.io, Inc.

package param

// GetHypervisorTypesDetailParam GetHypervisorTypes detail param
type GetHypervisorTypesDetailParam struct {
}

// GetHypervisorTypesParam GetHypervisorTypes request param
type GetHypervisorTypesParam struct {
	BaseParam
	Params GetHypervisorTypesDetailParam `json:"params"`
}
