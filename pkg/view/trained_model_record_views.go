// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TrainedModelRecordInventoryView TrainedModelRecord
type TrainedModelRecordInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"modelUuid,omitempty"`
	rest string `json:"sourceModelUuid,omitempty"`
	rest string `json:"modelServiceInstanceGroupUuid,omitempty"`
	rest string `json:"datasetUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

