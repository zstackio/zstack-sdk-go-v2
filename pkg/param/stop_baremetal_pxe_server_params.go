// Copyright (c) ZStack.io, Inc.

package param

// StopBaremetalPxeServerDetailParam StopBaremetalPxeServer detail param
type StopBaremetalPxeServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StopBaremetalPxeServerParam StopBaremetalPxeServer request param
type StopBaremetalPxeServerParam struct {
	BaseParam
	Params StopBaremetalPxeServerDetailParam `json:"params"`
}
