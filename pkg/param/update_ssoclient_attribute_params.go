// Copyright (c) ZStack.io, Inc.

package param

// UpdateSSOClientAttributeDetailParam UpdateSSOClientAttribute detail param
type UpdateSSOClientAttributeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateSSOClientAttributeParam UpdateSSOClientAttribute request param
type UpdateSSOClientAttributeParam struct {
	BaseParam
	Params UpdateSSOClientAttributeDetailParam `json:"params"`
}
