// Copyright (c) ZStack.io, Inc.

package param

// UpdateBaremetalInstanceDetailParam UpdateBaremetalInstance detail param
type UpdateBaremetalInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Password string `json:"password,omitempty"`
	Platform string `json:"platform,omitempty"`
}

// UpdateBaremetalInstanceParam UpdateBaremetalInstance request param
type UpdateBaremetalInstanceParam struct {
	BaseParam
	Params UpdateBaremetalInstanceDetailParam `json:"params"`
}
