// Copyright (c) ZStack.io, Inc.

package param

// UpdateVCenterDetailParam UpdateVCenter detail param
type UpdateVCenterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DomainName string `json:"domainName,omitempty"`
	Port int `json:"port,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateVCenterParam UpdateVCenter request param
type UpdateVCenterParam struct {
	BaseParam
	Params UpdateVCenterDetailParam `json:"params"`
}
