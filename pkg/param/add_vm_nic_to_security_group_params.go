// Copyright (c) ZStack.io, Inc.

package param

// AddVmNicToSecurityGroupDetailParam AddVmNicToSecurityGroup detail param
type AddVmNicToSecurityGroupDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// AddVmNicToSecurityGroupParam AddVmNicToSecurityGroup request param
type AddVmNicToSecurityGroupParam struct {
	BaseParam
	Params AddVmNicToSecurityGroupDetailParam `json:"params"`
}
