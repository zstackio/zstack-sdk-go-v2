// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelCenterCapacityInventoryView ModelCenterCapacity
type ModelCenterCapacityInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int64 `json:"modelUsedCapacity,omitempty"`
	rest int64 `json:"modelServiceUsedCapacity,omitempty"`
	rest int64 `json:"datasetUsedCapacity,omitempty"`
	rest int64 `json:"fineTuneUsedCapacity,omitempty"`
	rest int64 `json:"modelEvaluationUsedCapacity,omitempty"`
	rest int64 `json:"temporaryUsedCapacity,omitempty"`
	rest int64 `json:"cacheUsedCapacity,omitempty"`
}

