// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateIAM2ProjectRoleParamDetail CreateIAM2ProjectRole detail param
type CreateIAM2ProjectRoleParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Statements []PolicyStatementParam `json:"statements,omitempty"`
	PolicyUuids []string `json:"policyUuids,omitempty"`
	Identity *string `json:"identity,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectRoleParam CreateIAM2ProjectRole request param
type CreateIAM2ProjectRoleParam struct {
	BaseParam
	Params CreateIAM2ProjectRoleParamDetail `json:"params"`
}
