// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunProxyVSwitchInventoryView AliyunProxyVSwitch
type AliyunProxyVSwitchInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AliyunProxyVpcUuid string `json:"aliyunProxyVpcUuid,omitempty"`
	VpcL3NetworkUuid string `json:"vpcL3NetworkUuid,omitempty"`
	Status string `json:"status,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

