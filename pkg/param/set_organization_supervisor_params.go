// Copyright (c) ZStack.io, Inc.

package param

// SetOrganizationSupervisorDetailParam SetOrganizationSupervisor detail param
type SetOrganizationSupervisorDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
}

// SetOrganizationSupervisorParam SetOrganizationSupervisor request param
type SetOrganizationSupervisorParam struct {
	BaseParam
	Params SetOrganizationSupervisorDetailParam `json:"params"`
}
