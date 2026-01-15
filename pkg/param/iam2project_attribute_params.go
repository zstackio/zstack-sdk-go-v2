// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateIAM2ProjectAttributeParamDetail UpdateIAM2ProjectAttribute detail param
type UpdateIAM2ProjectAttributeParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2ProjectAttributeParam UpdateIAM2ProjectAttribute request param
type UpdateIAM2ProjectAttributeParam struct {
	BaseParam
	Params UpdateIAM2ProjectAttributeParamDetail `json:"updateIAM2ProjectAttribute"`
}
