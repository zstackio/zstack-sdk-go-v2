// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateIAM2OrganizationAttributeParamDetail UpdateIAM2OrganizationAttribute detail param
type UpdateIAM2OrganizationAttributeParamDetail struct {
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2OrganizationAttributeParam UpdateIAM2OrganizationAttribute request param
type UpdateIAM2OrganizationAttributeParam struct {
	BaseParam
	Params UpdateIAM2OrganizationAttributeParamDetail `json:"updateIAM2OrganizationAttribute"`
}
