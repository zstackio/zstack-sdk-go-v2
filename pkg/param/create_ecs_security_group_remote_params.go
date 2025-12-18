// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsSecurityGroupRemoteDetailParam CreateEcsSecurityGroupRemote detail param
type CreateEcsSecurityGroupRemoteDetailParam struct {
	VpcUuid string `json:"vpcUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	Strategy string `json:"strategy,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsSecurityGroupRemoteParam CreateEcsSecurityGroupRemote request param
type CreateEcsSecurityGroupRemoteParam struct {
	BaseParam
	Params CreateEcsSecurityGroupRemoteDetailParam `json:"params"`
}
