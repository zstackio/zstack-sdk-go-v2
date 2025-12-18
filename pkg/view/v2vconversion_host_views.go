// Copyright (c) ZStack.io, Inc.

package view

import "time"

// V2VConversionHostInventoryView V2VConversionHost
type V2VConversionHostInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"storagePath,omitempty"`
	rest string `json:"state,omitempty"`
	rest int64 `json:"totalSize,omitempty"`
	rest int64 `json:"availableSize,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"hostStatus,omitempty"`
	rest string `json:"hostState,omitempty"`
}

