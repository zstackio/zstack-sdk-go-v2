// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunProxyVSwitchInventoryView AliyunProxyVSwitch
type AliyunProxyVSwitchInventoryView struct {
	BaseInfoView
	BaseTimeView
	AliyunProxyVpcUuid *string `json:"aliyunProxyVpcUuid,omitempty"`
	VpcL3NetworkUuid *string `json:"vpcL3NetworkUuid,omitempty"`
	Status *string `json:"status,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// UpdateAliyunProxyVSwitchEventView UpdateAliyunProxyVSwitchEvent
type UpdateAliyunProxyVSwitchEventView struct {
	Inventory AliyunProxyVSwitchInventoryView `json:"inventory,omitempty"`
}

// QueryAliyunProxyVSwitchView QueryAliyunProxyVSwitch
type QueryAliyunProxyVSwitchView struct {
	Inventories []AliyunProxyVSwitchInventoryView `json:"inventories,omitempty"`
}

// CreateAliyunProxyVSwitchEventView CreateAliyunProxyVSwitchEvent
type CreateAliyunProxyVSwitchEventView struct {
	Inventory AliyunProxyVSwitchInventoryView `json:"inventory,omitempty"`
}

// DeleteAliyunProxyVSwitchEventView DeleteAliyunProxyVSwitchEvent
type DeleteAliyunProxyVSwitchEventView struct {
	Success bool `json:"success,omitempty"`
}

