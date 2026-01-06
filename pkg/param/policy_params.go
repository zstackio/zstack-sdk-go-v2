// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeletePolicyParamDetail DeletePolicy detail param
type DeletePolicyParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePolicyParam DeletePolicy request param
type DeletePolicyParam struct {
	BaseParam
	Params DeletePolicyParamDetail `json:"params"`
}
// CreatePolicyParamDetail CreatePolicy detail param
type CreatePolicyParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Statements []PolicyStatementParam `json:"statements" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyParam CreatePolicy request param
type CreatePolicyParam struct {
	BaseParam
	Params CreatePolicyParamDetail `json:"params"`
}
