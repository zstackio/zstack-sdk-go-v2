// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateIAM2OrganizationAttributeParamDetail UpdateIAM2OrganizationAttribute detail param
type UpdateIAM2OrganizationAttributeParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2OrganizationAttributeParam UpdateIAM2OrganizationAttribute request param
type UpdateIAM2OrganizationAttributeParam struct {
	BaseParam
	UpdateIAM2OrganizationAttribute UpdateIAM2OrganizationAttributeParamDetail `json:"updateIAM2OrganizationAttribute"`
}
