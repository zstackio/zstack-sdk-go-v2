// Copyright (c) ZStack.io, Inc.

package param

// DestroyBaremetalInstanceDetailParam DestroyBaremetalInstance detail param
type DestroyBaremetalInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DestroyBaremetalInstanceParam DestroyBaremetalInstance request param
type DestroyBaremetalInstanceParam struct {
	BaseParam
	Params DestroyBaremetalInstanceDetailParam `json:"params"`
}
