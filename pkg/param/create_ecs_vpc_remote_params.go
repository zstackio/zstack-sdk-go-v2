// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsVpcRemoteDetailParam CreateEcsVpcRemote detail param
type CreateEcsVpcRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	CidrBlock string `json:"cidrBlock" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VRouterName string `json:"vRouterName" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsVpcRemoteParam CreateEcsVpcRemote request param
type CreateEcsVpcRemoteParam struct {
	BaseParam
	Params CreateEcsVpcRemoteDetailParam `json:"params"`
}
