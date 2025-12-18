// Copyright (c) ZStack.io, Inc.

package param

// AddRolesToIAM2VirtualIDDetailParam AddRolesToIAM2VirtualID detail param
type AddRolesToIAM2VirtualIDDetailParam struct {
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
	RoleUuids []string `json:"roleUuids" validate:"required"`
	ProjectUuid string `json:"projectUuid,omitempty"`
}

// AddRolesToIAM2VirtualIDParam AddRolesToIAM2VirtualID request param
type AddRolesToIAM2VirtualIDParam struct {
	BaseParam
	Params AddRolesToIAM2VirtualIDDetailParam `json:"params"`
}
