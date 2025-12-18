// Copyright (c) ZStack.io, Inc.

package param

// DetachSshKeyPairFromVmInstanceDetailParam DetachSshKeyPairFromVmInstance详细参数
type DetachSshKeyPairFromVmInstanceDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"sshKeyPairUuid" validate:"required"` // 必填
}

// DetachSshKeyPairFromVmInstanceParam DetachSshKeyPairFromVmInstance请求参数
type DetachSshKeyPairFromVmInstanceParam struct {
	BaseParam
	Params DetachSshKeyPairFromVmInstanceDetailParam `json:"params"` // 详细参数
}

