// Copyright (c) ZStack.io, Inc.

package param

// SetVmSshKeyDetailParam SetVmSshKey detail param
type SetVmSshKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SshKey string `json:"SshKey" validate:"required"`
}

// SetVmSshKeyParam SetVmSshKey request param
type SetVmSshKeyParam struct {
	BaseParam
	Params SetVmSshKeyDetailParam `json:"params"`
}
