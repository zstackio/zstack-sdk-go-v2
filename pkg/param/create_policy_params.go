// Copyright (c) ZStack.io, Inc.

package param

// CreatePolicyDetailParam CreatePolicy detail param
type CreatePolicyDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Statements []PolicyStatementParam `json:"statements" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyParam CreatePolicy request param
type CreatePolicyParam struct {
	BaseParam
	Params CreatePolicyDetailParam `json:"params"`
}
