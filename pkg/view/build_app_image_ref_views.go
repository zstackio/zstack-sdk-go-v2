// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BuildAppImageRefInventoryView BuildAppImageRef
type BuildAppImageRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"buildAppUuid,omitempty"`
	rest string `json:"imageName,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

