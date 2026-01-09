// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceGroupDatasetRefInventoryView ModelServiceGroupDatasetRef
type ModelServiceGroupDatasetRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	DatasetUuid *string `json:"datasetUuid,omitempty"`
	ModelServiceInstanceGroupUuid *string `json:"modelServiceInstanceGroupUuid,omitempty"`
}

