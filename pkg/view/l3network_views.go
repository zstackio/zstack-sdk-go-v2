// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L3NetworkInventoryView L3Network
type L3NetworkInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	InternalId int `json:"internalId,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
	State string `json:"state,omitempty"`
	DnsDomain string `json:"dnsDomain,omitempty"`
	System bool `json:"system,omitempty"`
	Category string `json:"category,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	EnableIPAM bool `json:"enableIPAM,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Dns []string `json:"dns,omitempty"`
	IpRanges []IpRangeInventoryView `json:"ipRanges,omitempty"`
	NetworkServices []NetworkServiceL3NetworkRefInventoryView `json:"networkServices,omitempty"`
	HostRoute []L3NetworkHostRouteInventoryView `json:"hostRoute,omitempty"`
	ReservedIpRanges []ReservedIpRangeInventoryView `json:"reservedIpRanges,omitempty"`
}

