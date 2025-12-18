// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectsOfVirtualIDDetailParam GetIAM2ProjectsOfVirtualID detail param
type GetIAM2ProjectsOfVirtualIDDetailParam struct {
}

// GetIAM2ProjectsOfVirtualIDParam GetIAM2ProjectsOfVirtualID request param
type GetIAM2ProjectsOfVirtualIDParam struct {
	BaseParam
	Params GetIAM2ProjectsOfVirtualIDDetailParam `json:"params"`
}
