// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectRoleDetailParam CreateIAM2ProjectRole detail param
type CreateIAM2ProjectRoleDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Statements []interface{} `json:"statements,omitempty"`
	PolicyUuids []string `json:"policyUuids,omitempty"`
	Identity string `json:"identity,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectRoleParam CreateIAM2ProjectRole request param
type CreateIAM2ProjectRoleParam struct {
	BaseParam
	Params CreateIAM2ProjectRoleDetailParam `json:"params"`
}
