// Copyright (c) ZStack.io, Inc.

package view

import "time"

// L3NetworkInventoryView L3Network
type L3NetworkInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest int `json:"internalId,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"l2NetworkUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"dnsDomain,omitempty"`
	rest bool `json:"system,omitempty"`
	rest string `json:"category,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest bool `json:"enableIPAM,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"dns,omitempty"`
	rest []IpRangeInventoryView `json:"ipRanges,omitempty"`
	rest []NetworkServiceL3NetworkRefInventoryView `json:"networkServices,omitempty"`
	rest []L3NetworkHostRouteInventoryView `json:"hostRoute,omitempty"`
	rest []ReservedIpRangeInventoryView `json:"reservedIpRanges,omitempty"`
}

