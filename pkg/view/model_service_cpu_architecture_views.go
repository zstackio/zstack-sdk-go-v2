// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceCpuArchitectureInventoryView ModelServiceCpuArchitecture
type ModelServiceCpuArchitectureInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ModelServiceUuid *string `json:"modelServiceUuid,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
}

