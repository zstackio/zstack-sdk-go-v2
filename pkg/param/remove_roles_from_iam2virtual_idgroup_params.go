// Copyright (c) ZStack.io, Inc.

package param

// RemoveRolesFromIAM2VirtualIDGroupDetailParam RemoveRolesFromIAM2VirtualIDGroup detail param
type RemoveRolesFromIAM2VirtualIDGroupDetailParam struct {
	RoleUuids []string `json:"roleUuids" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
	ProjectUuid string `json:"projectUuid,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDGroupParam RemoveRolesFromIAM2VirtualIDGroup request param
type RemoveRolesFromIAM2VirtualIDGroupParam struct {
	BaseParam
	Params RemoveRolesFromIAM2VirtualIDGroupDetailParam `json:"params"`
}
