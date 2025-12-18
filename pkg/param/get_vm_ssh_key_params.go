// Copyright (c) ZStack.io, Inc.

package param

// GetVmSshKeyDetailParam GetVmSshKey detail param
type GetVmSshKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmSshKeyParam GetVmSshKey request param
type GetVmSshKeyParam struct {
	BaseParam
	Params GetVmSshKeyDetailParam `json:"params"`
}
