// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunProxyVpcDetailParam DeleteAliyunProxyVpc detail param
type DeleteAliyunProxyVpcDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunProxyVpcParam DeleteAliyunProxyVpc request param
type DeleteAliyunProxyVpcParam struct {
	BaseParam
	Params DeleteAliyunProxyVpcDetailParam `json:"params"`
}
