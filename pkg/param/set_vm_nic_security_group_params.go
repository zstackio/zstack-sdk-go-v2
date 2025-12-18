// Copyright (c) ZStack.io, Inc.

package param

// SetVmNicSecurityGroupDetailParam SetVmNicSecurityGroup detail param
type SetVmNicSecurityGroupDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	Refs []interface{} `json:"refs" validate:"required"`
}

// SetVmNicSecurityGroupParam SetVmNicSecurityGroup request param
type SetVmNicSecurityGroupParam struct {
	BaseParam
	Params SetVmNicSecurityGroupDetailParam `json:"params"`
}
