// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateIAM2VirtualIDAttributeParamDetail UpdateIAM2VirtualIDAttribute detail param
type UpdateIAM2VirtualIDAttributeParamDetail struct {
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2VirtualIDAttributeParam UpdateIAM2VirtualIDAttribute request param
type UpdateIAM2VirtualIDAttributeParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDAttributeParamDetail `json:"updateIAM2VirtualIDAttribute"`
}
