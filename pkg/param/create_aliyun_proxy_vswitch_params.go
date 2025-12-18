// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunProxyVSwitchDetailParam CreateAliyunProxyVSwitch detail param
type CreateAliyunProxyVSwitchDetailParam struct {
	AliyunProxyVpcUuid string `json:"aliyunProxyVpcUuid" validate:"required"`
	VpcL3NetworkUuid string `json:"vpcL3NetworkUuid" validate:"required"`
	IsDefault bool `json:"isDefault" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunProxyVSwitchParam CreateAliyunProxyVSwitch request param
type CreateAliyunProxyVSwitchParam struct {
	BaseParam
	Params CreateAliyunProxyVSwitchDetailParam `json:"params"`
}
