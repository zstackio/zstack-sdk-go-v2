// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateAliyunProxyVSwitchParamDetail UpdateAliyunProxyVSwitch detail param
type UpdateAliyunProxyVSwitchParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Status string `json:"status,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// UpdateAliyunProxyVSwitchParam UpdateAliyunProxyVSwitch request param
type UpdateAliyunProxyVSwitchParam struct {
	BaseParam
	UpdateAliyunProxyVSwitch UpdateAliyunProxyVSwitchParamDetail `json:"updateAliyunProxyVSwitch"`
}
// CreateAliyunProxyVSwitchParamDetail CreateAliyunProxyVSwitch detail param
type CreateAliyunProxyVSwitchParamDetail struct {
	AliyunProxyVpcUuid string `json:"aliyunProxyVpcUuid" validate:"required"`
	VpcL3NetworkUuid string `json:"vpcL3NetworkUuid" validate:"required"`
	IsDefault bool `json:"isDefault" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunProxyVSwitchParam CreateAliyunProxyVSwitch request param
type CreateAliyunProxyVSwitchParam struct {
	BaseParam
	CreateAliyunProxyVSwitch CreateAliyunProxyVSwitchParamDetail `json:"createAliyunProxyVSwitch"`
}
// DeleteAliyunProxyVSwitchParamDetail DeleteAliyunProxyVSwitch detail param
type DeleteAliyunProxyVSwitchParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunProxyVSwitchParam DeleteAliyunProxyVSwitch request param
type DeleteAliyunProxyVSwitchParam struct {
	BaseParam
	DeleteAliyunProxyVSwitch DeleteAliyunProxyVSwitchParamDetail `json:"deleteAliyunProxyVSwitch"`
}
