// Copyright (c) ZStack.io, Inc.

package param

// SetVmSshKeyDetailParam SetVmSshKey详细参数
type SetVmSshKeyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"SshKey" validate:"required"` // 必填
}

// SetVmSshKeyParam SetVmSshKey请求参数
type SetVmSshKeyParam struct {
	BaseParam
	Params SetVmSshKeyDetailParam `json:"params"` // 详细参数
}

