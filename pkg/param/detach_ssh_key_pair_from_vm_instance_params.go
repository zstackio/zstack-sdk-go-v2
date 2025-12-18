// Copyright (c) ZStack.io, Inc.

package param

// DetachSshKeyPairFromVmInstanceDetailParam DetachSshKeyPairFromVmInstance detail param
type DetachSshKeyPairFromVmInstanceDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	SshKeyPairUuid string `json:"sshKeyPairUuid" validate:"required"`
}

// DetachSshKeyPairFromVmInstanceParam DetachSshKeyPairFromVmInstance request param
type DetachSshKeyPairFromVmInstanceParam struct {
	BaseParam
	Params DetachSshKeyPairFromVmInstanceDetailParam `json:"params"`
}
