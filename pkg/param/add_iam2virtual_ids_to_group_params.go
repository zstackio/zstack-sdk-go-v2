// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToGroupDetailParam AddIAM2VirtualIDsToGroup detail param
type AddIAM2VirtualIDsToGroupDetailParam struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// AddIAM2VirtualIDsToGroupParam AddIAM2VirtualIDsToGroup request param
type AddIAM2VirtualIDsToGroupParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToGroupDetailParam `json:"params"`
}
