// Copyright (c) ZStack.io, Inc.

package param

// AttachSecurityGroupToL3NetworkDetailParam AttachSecurityGroupToL3Network detail param
type AttachSecurityGroupToL3NetworkDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
}

// AttachSecurityGroupToL3NetworkParam AttachSecurityGroupToL3Network request param
type AttachSecurityGroupToL3NetworkParam struct {
	BaseParam
	Params AttachSecurityGroupToL3NetworkDetailParam `json:"params"`
}
