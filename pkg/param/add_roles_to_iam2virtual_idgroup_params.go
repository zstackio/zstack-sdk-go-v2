// Copyright (c) ZStack.io, Inc.

package param

// AddRolesToIAM2VirtualIDGroupDetailParam AddRolesToIAM2VirtualIDGroup detail param
type AddRolesToIAM2VirtualIDGroupDetailParam struct {
	RoleUuids []string `json:"roleUuids" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
	ProjectUuid string `json:"projectUuid,omitempty"`
}

// AddRolesToIAM2VirtualIDGroupParam AddRolesToIAM2VirtualIDGroup request param
type AddRolesToIAM2VirtualIDGroupParam struct {
	BaseParam
	Params AddRolesToIAM2VirtualIDGroupDetailParam `json:"params"`
}
