// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2NetworkDataView L2NetworkData
type L2NetworkDataView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	PoolUuid string `json:"poolUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Description string `json:"description,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	Vni string `json:"vni,omitempty"`
	Vlan int `json:"vlan,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

