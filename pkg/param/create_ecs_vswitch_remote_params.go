// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsVSwitchRemoteDetailParam CreateEcsVSwitchRemote detail param
type CreateEcsVSwitchRemoteDetailParam struct {
	VpcUuid string `json:"vpcUuid" validate:"required"`
	IdentityZoneUuid string `json:"identityZoneUuid" validate:"required"`
	CidrBlock string `json:"cidrBlock" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsVSwitchRemoteParam CreateEcsVSwitchRemote request param
type CreateEcsVSwitchRemoteParam struct {
	BaseParam
	Params CreateEcsVSwitchRemoteDetailParam `json:"params"`
}
