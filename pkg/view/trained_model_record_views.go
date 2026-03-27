// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TrainedModelRecordInventoryView TrainedModelRecord
type TrainedModelRecordInventoryView struct {
	BaseInfoView
	BaseTimeView
	ModelUuid string `json:"modelUuid,omitempty"`
	SourceModelUuid string `json:"sourceModelUuid,omitempty"`
	ModelServiceInstanceGroupUuid string `json:"modelServiceInstanceGroupUuid,omitempty"`
	DatasetUuid string `json:"datasetUuid,omitempty"`
}

// QueryTrainedModelRecordView QueryTrainedModelRecord
type QueryTrainedModelRecordView struct {
	Inventories []TrainedModelRecordInventoryView `json:"inventories,omitempty"`
}

