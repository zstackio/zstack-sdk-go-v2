// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DatasetInventoryView Dataset
type DatasetInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	ModelCenterUuid string `json:"modelCenterUuid,omitempty"`
	Size int64 `json:"size,omitempty"`
	System bool `json:"system,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

