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
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryDatasetView QueryDataset
type QueryDatasetView struct {
	Inventories []DatasetInventoryView `json:"inventories,omitempty"`
}

// CreateDatasetEventView CreateDatasetEvent
type CreateDatasetEventView struct {
	Inventory DatasetInventoryView `json:"inventory,omitempty"`
}

// DeleteDatasetEventView DeleteDatasetEvent
type DeleteDatasetEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateDatasetEventView UpdateDatasetEvent
type UpdateDatasetEventView struct {
	Inventory DatasetInventoryView `json:"inventory,omitempty"`
}

