// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromGroupDetailParam RemoveIAM2VirtualIDsFromGroup detail param
type RemoveIAM2VirtualIDsFromGroupDetailParam struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// RemoveIAM2VirtualIDsFromGroupParam RemoveIAM2VirtualIDsFromGroup request param
type RemoveIAM2VirtualIDsFromGroupParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromGroupDetailParam `json:"params"`
}
