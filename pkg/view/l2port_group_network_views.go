// Copyright (c) ZStack.io, Inc.

package view

import "time"

// L2PortGroupNetworkInventoryView L2PortGroupNetwork
type L2PortGroupNetworkInventoryView struct {
	rest string `json:"vSwitchUuid,omitempty"`
	rest string `json:"vlanMode,omitempty"`
	rest int `json:"vlanId,omitempty"`
	rest string `json:"vlanRanges,omitempty"`
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

