// Copyright (c) ZStack.io, Inc.

package param

// CreateRoleDetailParam CreateRole detail param
type CreateRoleDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Statements []interface{} `json:"statements,omitempty"`
	PolicyUuids []string `json:"policyUuids,omitempty"`
	Identity string `json:"identity,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRoleParam CreateRole request param
type CreateRoleParam struct {
	BaseParam
	Params CreateRoleDetailParam `json:"params"`
}
