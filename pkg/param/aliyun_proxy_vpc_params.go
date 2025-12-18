// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunProxyVpcDetailParam UpdateAliyunProxyVpc详细参数
type UpdateAliyunProxyVpcDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest bool `json:"isDefault,omitempty"`
	rest string `json:"status,omitempty"`
}

// UpdateAliyunProxyVpcParam UpdateAliyunProxyVpc请求参数
type UpdateAliyunProxyVpcParam struct {
	BaseParam
	Params UpdateAliyunProxyVpcDetailParam `json:"params"` // 详细参数
}

