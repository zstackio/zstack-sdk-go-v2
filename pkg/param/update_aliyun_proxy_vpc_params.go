// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunProxyVpcDetailParam UpdateAliyunProxyVpc detail param
type UpdateAliyunProxyVpcDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Status string `json:"status,omitempty"`
}

// UpdateAliyunProxyVpcParam UpdateAliyunProxyVpc request param
type UpdateAliyunProxyVpcParam struct {
	BaseParam
	Params UpdateAliyunProxyVpcDetailParam `json:"params"`
}
