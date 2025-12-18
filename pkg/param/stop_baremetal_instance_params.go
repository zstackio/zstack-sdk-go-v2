// Copyright (c) ZStack.io, Inc.

package param

// StopBaremetalInstanceDetailParam StopBaremetalInstance detail param
type StopBaremetalInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Type string `json:"type,omitempty"`
}

// StopBaremetalInstanceParam StopBaremetalInstance request param
type StopBaremetalInstanceParam struct {
	BaseParam
	Params StopBaremetalInstanceDetailParam `json:"params"`
}
