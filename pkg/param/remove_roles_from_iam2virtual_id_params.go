// Copyright (c) ZStack.io, Inc.

package param

// RemoveRolesFromIAM2VirtualIDDetailParam RemoveRolesFromIAM2VirtualID detail param
type RemoveRolesFromIAM2VirtualIDDetailParam struct {
	RoleUuids []string `json:"roleUuids" validate:"required"`
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
	ProjectUuid string `json:"projectUuid,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDParam RemoveRolesFromIAM2VirtualID request param
type RemoveRolesFromIAM2VirtualIDParam struct {
	BaseParam
	Params RemoveRolesFromIAM2VirtualIDDetailParam `json:"params"`
}
