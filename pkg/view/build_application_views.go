// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BuildApplicationInventoryView BuildApplication
type BuildApplicationInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"buildSystemUuid,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"appMetaData,omitempty"`
	rest string `json:"appId,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

