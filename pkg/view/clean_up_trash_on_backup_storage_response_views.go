// Copyright (c) ZStack.io, Inc.

package view

// CleanUpTrashOnBackupStorageEventView CleanUpTrashOnBackupStorageEvent
type CleanUpTrashOnBackupStorageEventView struct {
	Result CleanTrashResultView `json:"result,omitempty"`
	Results []TrashCleanupResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

