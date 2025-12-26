// Copyright (c) ZStack.io, Inc.

package view

// CleanUpTrashOnPrimaryStorageEventView CleanUpTrashOnPrimaryStorageEvent
type CleanUpTrashOnPrimaryStorageEventView struct {
	Result CleanTrashResultView `json:"result,omitempty"`
	Results []TrashCleanupResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

