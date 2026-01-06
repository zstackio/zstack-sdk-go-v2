// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DataCenterInventoryView DataCenter
type DataCenterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Deleted string `json:"deleted,omitempty"`
	RegionName string `json:"regionName,omitempty"`
	DcType string `json:"dcType,omitempty"`
	RegionId string `json:"regionId,omitempty"`
	Description string `json:"description,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// AddDataCenterFromRemoteEventView AddDataCenterFromRemoteEvent
type AddDataCenterFromRemoteEventView struct {
	Inventory DataCenterInventoryView `json:"inventory,omitempty"`
}

// QueryDataCenterFromLocalView QueryDataCenterFromLocal
type QueryDataCenterFromLocalView struct {
	Inventories []DataCenterInventoryView `json:"inventories,omitempty"`
}

