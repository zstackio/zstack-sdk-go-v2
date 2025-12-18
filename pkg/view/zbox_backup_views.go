// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ZBoxBackupInventoryView ZBoxBackup
type ZBoxBackupInventoryView struct {
	rest string `json:"zBoxUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest int64 `json:"totalSize,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

