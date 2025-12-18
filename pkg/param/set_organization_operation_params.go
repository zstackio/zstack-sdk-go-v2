// Copyright (c) ZStack.io, Inc.

package param

// SetOrganizationOperationDetailParam SetOrganizationOperation detail param
type SetOrganizationOperationDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
}

// SetOrganizationOperationParam SetOrganizationOperation request param
type SetOrganizationOperationParam struct {
	BaseParam
	Params SetOrganizationOperationDetailParam `json:"params"`
}
