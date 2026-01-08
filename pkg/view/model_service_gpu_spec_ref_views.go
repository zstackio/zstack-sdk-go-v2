// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceGpuSpecRefInventoryView ModelServiceGpuSpecRef
type ModelServiceGpuSpecRefInventoryView struct {
	Id       int64  `json:"id,omitempty"`
	RefUuid  int64  `json:"refUuid,omitempty"`
	SpecUuid string `json:"specUuid,omitempty"`
}
