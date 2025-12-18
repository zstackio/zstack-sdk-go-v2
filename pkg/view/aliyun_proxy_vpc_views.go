// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunProxyVpcInventoryView AliyunProxyVpc
type AliyunProxyVpcInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vpcName,omitempty"`
	rest string `json:"cidrBlock,omitempty"`
	rest string `json:"vRouterUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest []AliyunProxyVSwitchInventoryView `json:"aliyunProxyVSwitches,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest bool `json:"isDefault,omitempty"`
}

