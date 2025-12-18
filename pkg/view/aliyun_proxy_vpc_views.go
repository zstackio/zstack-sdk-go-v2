// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunProxyVpcInventoryView AliyunProxyVpc
type AliyunProxyVpcInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VpcName string `json:"vpcName,omitempty"`
	CidrBlock string `json:"cidrBlock,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	Status string `json:"status,omitempty"`
	AliyunProxyVSwitches []AliyunProxyVSwitchInventoryView `json:"aliyunProxyVSwitches,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

