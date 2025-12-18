// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmNicFromSecurityGroupDetailParam DeleteVmNicFromSecurityGroup详细参数
type DeleteVmNicFromSecurityGroupDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest []string `json:"vmNicUuids" validate:"required"` // 必填
}

// DeleteVmNicFromSecurityGroupParam DeleteVmNicFromSecurityGroup请求参数
type DeleteVmNicFromSecurityGroupParam struct {
	BaseParam
	Params DeleteVmNicFromSecurityGroupDetailParam `json:"params"` // 详细参数
}

