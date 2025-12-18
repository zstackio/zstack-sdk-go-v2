// Copyright (c) ZStack.io, Inc.

package param

// GetVmSshKeyDetailParam GetVmSshKey详细参数
type GetVmSshKeyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmSshKeyParam GetVmSshKey请求参数
type GetVmSshKeyParam struct {
	BaseParam
	Params GetVmSshKeyDetailParam `json:"params"` // 详细参数
}

