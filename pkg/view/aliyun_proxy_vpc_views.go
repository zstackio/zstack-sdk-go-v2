// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunProxyVpcInventoryView AliyunProxyVpc
type AliyunProxyVpcInventoryView struct {
	BaseInfoView
	BaseTimeView
	VpcName string `json:"vpcName,omitempty"`
	CidrBlock string `json:"cidrBlock,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	Status string `json:"status,omitempty"`
	AliyunProxyVSwitches []AliyunProxyVSwitchInventoryView `json:"aliyunProxyVSwitches,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// CreateAliyunProxyVpcEventView CreateAliyunProxyVpcEvent
type CreateAliyunProxyVpcEventView struct {
	Inventory AliyunProxyVpcInventoryView `json:"inventory,omitempty"`
}

// UpdateAliyunProxyVpcEventView UpdateAliyunProxyVpcEvent
type UpdateAliyunProxyVpcEventView struct {
	Inventory AliyunProxyVpcInventoryView `json:"inventory,omitempty"`
}

// QueryAliyunProxyVpcView QueryAliyunProxyVpc
type QueryAliyunProxyVpcView struct {
	Inventories []AliyunProxyVpcInventoryView `json:"inventories,omitempty"`
}

// DeleteAliyunProxyVpcEventView DeleteAliyunProxyVpcEvent
type DeleteAliyunProxyVpcEventView struct {
	Success bool `json:"success,omitempty"`
}

