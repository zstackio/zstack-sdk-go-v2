// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunProxyVpcDetailParam CreateAliyunProxyVpc detail param
type CreateAliyunProxyVpcDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	CidrBlock string `json:"cidrBlock" validate:"required"`
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	IsDefault bool `json:"isDefault" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunProxyVpcParam CreateAliyunProxyVpc request param
type CreateAliyunProxyVpcParam struct {
	BaseParam
	Params CreateAliyunProxyVpcDetailParam `json:"params"`
}
