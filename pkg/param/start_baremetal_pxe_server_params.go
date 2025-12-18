// Copyright (c) ZStack.io, Inc.

package param

// StartBaremetalPxeServerDetailParam StartBaremetalPxeServer detail param
type StartBaremetalPxeServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StartBaremetalPxeServerParam StartBaremetalPxeServer request param
type StartBaremetalPxeServerParam struct {
	BaseParam
	Params StartBaremetalPxeServerDetailParam `json:"params"`
}
