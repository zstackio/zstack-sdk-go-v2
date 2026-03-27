// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateRoleParamDetail CreateRole detail param
type CreateRoleParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Statements []PolicyStatementParam `json:"statements,omitempty"`
	PolicyUuids []string `json:"policyUuids,omitempty"`
	Identity *string `json:"identity,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRoleParam CreateRole request param
type CreateRoleParam struct {
	BaseParam
	Params CreateRoleParamDetail `json:"params"`
}
// DeleteRoleParamDetail DeleteRole detail param
type DeleteRoleParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteRoleParam DeleteRole request param
type DeleteRoleParam struct {
	BaseParam
	Params DeleteRoleParamDetail `json:"deleteRole"`
}
// UpdateRoleParamDetail UpdateRole detail param
type UpdateRoleParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Statements []PolicyStatementParam `json:"statements,omitempty"`
	PolicyUuids []string `json:"policyUuids,omitempty"`
}

// UpdateRoleParam UpdateRole request param
type UpdateRoleParam struct {
	BaseParam
	Params UpdateRoleParamDetail `json:"updateRole"`
}
