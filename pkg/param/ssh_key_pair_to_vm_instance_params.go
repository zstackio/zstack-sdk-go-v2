// Copyright (c) ZStack.io, Inc.

package param

// AttachSshKeyPairToVmInstanceDetailParam AttachSshKeyPairToVmInstance详细参数
type AttachSshKeyPairToVmInstanceDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"sshKeyPairUuid" validate:"required"` // 必填
}

// AttachSshKeyPairToVmInstanceParam AttachSshKeyPairToVmInstance请求参数
type AttachSshKeyPairToVmInstanceParam struct {
	BaseParam
	Params AttachSshKeyPairToVmInstanceDetailParam `json:"params"` // 详细参数
}

