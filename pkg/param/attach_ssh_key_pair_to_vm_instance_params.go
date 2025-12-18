// Copyright (c) ZStack.io, Inc.

package param

// AttachSshKeyPairToVmInstanceDetailParam AttachSshKeyPairToVmInstance detail param
type AttachSshKeyPairToVmInstanceDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	SshKeyPairUuid string `json:"sshKeyPairUuid" validate:"required"`
}

// AttachSshKeyPairToVmInstanceParam AttachSshKeyPairToVmInstance request param
type AttachSshKeyPairToVmInstanceParam struct {
	BaseParam
	Params AttachSshKeyPairToVmInstanceDetailParam `json:"params"`
}
