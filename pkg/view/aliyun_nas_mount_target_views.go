// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunNasMountTargetInventoryView AliyunNasMountTarget
type AliyunNasMountTargetInventoryView struct {
	rest string `json:"accessGroupUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"mountDomain,omitempty"`
	rest string `json:"nasFileSystemUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

