// Copyright (c) ZStack.io, Inc.

package param

// SetVmSecurityLevelDetailParam SetVmSecurityLevel详细参数
type SetVmSecurityLevelDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"securityLevel,omitempty"`
}

// SetVmSecurityLevelParam SetVmSecurityLevel请求参数
type SetVmSecurityLevelParam struct {
	BaseParam
	Params SetVmSecurityLevelDetailParam `json:"params"` // 详细参数
}

