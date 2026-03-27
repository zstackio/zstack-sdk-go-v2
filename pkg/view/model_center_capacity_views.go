// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ModelCenterCapacityInventoryView ModelCenterCapacity
type ModelCenterCapacityInventoryView struct {
	BaseInfoView
	BaseTimeView
	ModelUsedCapacity int64 `json:"modelUsedCapacity,omitempty"`
	ModelServiceUsedCapacity int64 `json:"modelServiceUsedCapacity,omitempty"`
	DatasetUsedCapacity int64 `json:"datasetUsedCapacity,omitempty"`
	FineTuneUsedCapacity int64 `json:"fineTuneUsedCapacity,omitempty"`
	ModelEvaluationUsedCapacity int64 `json:"modelEvaluationUsedCapacity,omitempty"`
	TemporaryUsedCapacity int64 `json:"temporaryUsedCapacity,omitempty"`
	CacheUsedCapacity int64 `json:"cacheUsedCapacity,omitempty"`
}

