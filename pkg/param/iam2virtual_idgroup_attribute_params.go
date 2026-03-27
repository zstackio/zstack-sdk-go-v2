// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateIAM2VirtualIDGroupAttributeParamDetail UpdateIAM2VirtualIDGroupAttribute detail param
type UpdateIAM2VirtualIDGroupAttributeParamDetail struct {
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2VirtualIDGroupAttributeParam UpdateIAM2VirtualIDGroupAttribute request param
type UpdateIAM2VirtualIDGroupAttributeParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDGroupAttributeParamDetail `json:"updateIAM2VirtualIDGroupAttribute"`
}
