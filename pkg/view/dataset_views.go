// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DatasetInventoryView Dataset
type DatasetInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"modelCenterUuid,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest bool `json:"system,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

