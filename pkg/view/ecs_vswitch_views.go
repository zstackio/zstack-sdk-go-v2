// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsVSwitchInventoryView EcsVSwitch
type EcsVSwitchInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VSwitchId string `json:"vSwitchId,omitempty"`
	Status string `json:"status,omitempty"`
	CidrBlock string `json:"cidrBlock,omitempty"`
	AvailableIpAddressCount int `json:"availableIpAddressCount,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	EcsVpcUuid string `json:"ecsVpcUuid,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

