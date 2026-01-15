// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateSSOClientAttributeParamDetail UpdateSSOClientAttribute detail param
type UpdateSSOClientAttributeParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateSSOClientAttributeParam UpdateSSOClientAttribute request param
type UpdateSSOClientAttributeParam struct {
	BaseParam
	UpdateSSOClientAttribute UpdateSSOClientAttributeParamDetail `json:"updateSSOClientAttribute"`
}
