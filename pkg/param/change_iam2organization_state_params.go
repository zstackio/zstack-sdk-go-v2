// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2OrganizationStateDetailParam ChangeIAM2OrganizationState detail param
type ChangeIAM2OrganizationStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2OrganizationStateParam ChangeIAM2OrganizationState request param
type ChangeIAM2OrganizationStateParam struct {
	BaseParam
	Params ChangeIAM2OrganizationStateDetailParam `json:"params"`
}
