// Copyright (c) ZStack.io, Inc.

package param

// SetOrganizationSupervisorDetailParam SetOrganizationSupervisor详细参数
type SetOrganizationSupervisorDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"virtualIDUuid" validate:"required"` // 必填
}

// SetOrganizationSupervisorParam SetOrganizationSupervisor请求参数
type SetOrganizationSupervisorParam struct {
	BaseParam
	Params SetOrganizationSupervisorDetailParam `json:"params"` // 详细参数
}

