// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DirectoryInventoryView Directory
type DirectoryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"groupName,omitempty"`
	rest string `json:"parentUuid,omitempty"`
	rest string `json:"rootDirectoryUuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

