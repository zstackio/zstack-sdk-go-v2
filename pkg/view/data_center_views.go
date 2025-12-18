// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DataCenterInventoryView DataCenter
type DataCenterInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"deleted,omitempty"`
	rest string `json:"regionName,omitempty"`
	rest string `json:"dcType,omitempty"`
	rest string `json:"regionId,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"endpoint,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

