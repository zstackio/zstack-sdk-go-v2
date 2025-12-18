// Copyright (c) ZStack.io, Inc.

package param

// RebootBaremetalInstanceDetailParam RebootBaremetalInstance detail param
type RebootBaremetalInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PxeBoot bool `json:"pxeBoot,omitempty"`
}

// RebootBaremetalInstanceParam RebootBaremetalInstance request param
type RebootBaremetalInstanceParam struct {
	BaseParam
	Params RebootBaremetalInstanceDetailParam `json:"params"`
}
