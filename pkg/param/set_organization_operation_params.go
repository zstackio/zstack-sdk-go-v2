// Copyright (c) ZStack.io, Inc.

package param

// SetOrganizationOperationDetailParam SetOrganizationOperation详细参数
type SetOrganizationOperationDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"virtualIDUuid" validate:"required"` // 必填
}

// SetOrganizationOperationParam SetOrganizationOperation请求参数
type SetOrganizationOperationParam struct {
	BaseParam
	Params SetOrganizationOperationDetailParam `json:"params"` // 详细参数
}

