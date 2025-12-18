// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NasMountTargetInventoryView NasMountTarget
type NasMountTargetInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"mountDomain,omitempty"`
	rest string `json:"nasFileSystemUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

