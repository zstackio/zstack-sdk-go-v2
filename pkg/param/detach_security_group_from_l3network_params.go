// Copyright (c) ZStack.io, Inc.

package param

// DetachSecurityGroupFromL3NetworkDetailParam DetachSecurityGroupFromL3Network detail param
type DetachSecurityGroupFromL3NetworkDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
}

// DetachSecurityGroupFromL3NetworkParam DetachSecurityGroupFromL3Network request param
type DetachSecurityGroupFromL3NetworkParam struct {
	BaseParam
	Params DetachSecurityGroupFromL3NetworkDetailParam `json:"params"`
}
