// Copyright (c) ZStack.io, Inc.

package view

import "time"

// L2VxlanNetworkInventoryView L2VxlanNetwork
type L2VxlanNetworkInventoryView struct {
	rest int `json:"vni,omitempty"`
	rest string `json:"poolUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"physicalInterface,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest int `json:"virtualNetworkId,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest string `json:"pvlan,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"attachedClusterUuids,omitempty"`
}

