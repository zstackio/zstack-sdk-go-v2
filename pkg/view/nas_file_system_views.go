// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NasFileSystemInventoryView NasFileSystem
type NasFileSystemInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"fileSystemId,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

