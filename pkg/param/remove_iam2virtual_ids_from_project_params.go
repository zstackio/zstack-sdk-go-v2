// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromProjectDetailParam RemoveIAM2VirtualIDsFromProject detail param
type RemoveIAM2VirtualIDsFromProjectDetailParam struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// RemoveIAM2VirtualIDsFromProjectParam RemoveIAM2VirtualIDsFromProject request param
type RemoveIAM2VirtualIDsFromProjectParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromProjectDetailParam `json:"params"`
}
