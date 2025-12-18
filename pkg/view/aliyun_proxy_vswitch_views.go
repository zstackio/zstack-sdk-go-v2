// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunProxyVSwitchInventoryView AliyunProxyVSwitch
type AliyunProxyVSwitchInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"aliyunProxyVpcUuid,omitempty"`
	rest string `json:"vpcL3NetworkUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest bool `json:"isDefault,omitempty"`
}

