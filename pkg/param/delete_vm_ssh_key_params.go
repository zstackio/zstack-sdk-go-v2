// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmSshKeyDetailParam DeleteVmSshKey detail param
type DeleteVmSshKeyDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
}

// DeleteVmSshKeyParam DeleteVmSshKey request param
type DeleteVmSshKeyParam struct {
	BaseParam
	Params DeleteVmSshKeyDetailParam `json:"params"`
}
