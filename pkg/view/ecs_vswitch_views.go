// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EcsVSwitchInventoryView EcsVSwitch
type EcsVSwitchInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vSwitchId,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"cidrBlock,omitempty"`
	rest int `json:"availableIpAddressCount,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"ecsVpcUuid,omitempty"`
	rest string `json:"identityZoneUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

