// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceGpuVendorInventoryView ModelServiceGpuVendor
type ModelServiceGpuVendorInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	ModelServiceUuid *string `json:"modelServiceUuid,omitempty"`
	GpuVendor *string `json:"gpuVendor,omitempty"`
	SpecRefs []ModelServiceGpuSpecRefInventoryView `json:"specRefs,omitempty"`
}

