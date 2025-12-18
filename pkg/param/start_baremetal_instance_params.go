// Copyright (c) ZStack.io, Inc.

package param

// StartBaremetalInstanceDetailParam StartBaremetalInstance detail param
type StartBaremetalInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PxeBoot bool `json:"pxeBoot,omitempty"`
}

// StartBaremetalInstanceParam StartBaremetalInstance request param
type StartBaremetalInstanceParam struct {
	BaseParam
	Params StartBaremetalInstanceDetailParam `json:"params"`
}
