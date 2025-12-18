// Copyright (c) ZStack.io, Inc.

package param

// AddVmNicToSecurityGroupDetailParam AddVmNicToSecurityGroup详细参数
type AddVmNicToSecurityGroupDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest []string `json:"vmNicUuids" validate:"required"` // 必填
}

// AddVmNicToSecurityGroupParam AddVmNicToSecurityGroup请求参数
type AddVmNicToSecurityGroupParam struct {
	BaseParam
	Params AddVmNicToSecurityGroupDetailParam `json:"params"` // 详细参数
}

