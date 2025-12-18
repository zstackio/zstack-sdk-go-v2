// Copyright (c) ZStack.io, Inc.

package param

// SetSecurityMachineKeyDetailParam SetSecurityMachineKey详细参数
type SetSecurityMachineKeyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"tokenName" validate:"required"` // 必填
	rest bool `json:"dryRun,omitempty"`
}

// SetSecurityMachineKeyParam SetSecurityMachineKey请求参数
type SetSecurityMachineKeyParam struct {
	BaseParam
	Params SetSecurityMachineKeyDetailParam `json:"params"` // 详细参数
}

