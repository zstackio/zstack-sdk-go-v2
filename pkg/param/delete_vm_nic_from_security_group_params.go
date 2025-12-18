// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmNicFromSecurityGroupDetailParam DeleteVmNicFromSecurityGroup detail param
type DeleteVmNicFromSecurityGroupDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// DeleteVmNicFromSecurityGroupParam DeleteVmNicFromSecurityGroup request param
type DeleteVmNicFromSecurityGroupParam struct {
	BaseParam
	Params DeleteVmNicFromSecurityGroupDetailParam `json:"params"`
}
