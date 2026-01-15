// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateAliyunProxyVpcParamDetail CreateAliyunProxyVpc detail param
type CreateAliyunProxyVpcParamDetail struct {
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
	CreateAliyunProxyVpc CreateAliyunProxyVpcParamDetail `json:"createAliyunProxyVpc"`
}
// UpdateAliyunProxyVpcParamDetail UpdateAliyunProxyVpc detail param
type UpdateAliyunProxyVpcParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Status string `json:"status,omitempty"`
}

// UpdateAliyunProxyVpcParam UpdateAliyunProxyVpc request param
type UpdateAliyunProxyVpcParam struct {
	BaseParam
	UpdateAliyunProxyVpc UpdateAliyunProxyVpcParamDetail `json:"updateAliyunProxyVpc"`
}
// DeleteAliyunProxyVpcParamDetail DeleteAliyunProxyVpc detail param
type DeleteAliyunProxyVpcParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunProxyVpcParam DeleteAliyunProxyVpc request param
type DeleteAliyunProxyVpcParam struct {
	BaseParam
	DeleteAliyunProxyVpc DeleteAliyunProxyVpcParamDetail `json:"deleteAliyunProxyVpc"`
}
