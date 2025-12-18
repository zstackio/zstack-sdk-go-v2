// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectRoleDetailParam CreateIAM2ProjectRole详细参数
type CreateIAM2ProjectRoleDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest []interface{} `json:"statements,omitempty"`
	rest []string `json:"policyUuids,omitempty"`
	rest string `json:"identity,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectRoleParam CreateIAM2ProjectRole请求参数
type CreateIAM2ProjectRoleParam struct {
	BaseParam
	Params CreateIAM2ProjectRoleDetailParam `json:"params"` // 详细参数
}

