// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// SetDGpuProfileParamDetail SetDGpuProfile detail param
type SetDGpuProfileParamDetail struct {
	MemorySizes []int64 `json:"memorySizes" validate:"required"`
	ShmemSize *int64 `json:"shmemSize,omitempty"`
}

// SetDGpuProfileParam SetDGpuProfile request param
type SetDGpuProfileParam struct {
	BaseParam
	Params SetDGpuProfileParamDetail `json:"setDGpuProfile"`
}
