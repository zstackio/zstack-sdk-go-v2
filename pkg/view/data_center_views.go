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
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

