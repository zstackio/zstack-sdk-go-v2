// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelServiceGpuVendorInventoryView ModelServiceGpuVendor
type ModelServiceGpuVendorInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"modelServiceUuid,omitempty"`
	rest string `json:"gpuVendor,omitempty"`
	rest []ModelServiceGpuSpecRefInventoryView `json:"specRefs,omitempty"`
}

