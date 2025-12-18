// Copyright (c) ZStack.io, Inc.

package param

// SetVmNicSecurityGroupDetailParam SetVmNicSecurityGroup详细参数
type SetVmNicSecurityGroupDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest []interface{} `json:"refs" validate:"required"` // 必填
}

// SetVmNicSecurityGroupParam SetVmNicSecurityGroup请求参数
type SetVmNicSecurityGroupParam struct {
	BaseParam
	Params SetVmNicSecurityGroupDetailParam `json:"params"` // 详细参数
}

