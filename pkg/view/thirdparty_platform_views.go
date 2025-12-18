// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ThirdpartyPlatformInventoryView ThirdpartyPlatform
type ThirdpartyPlatformInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"template,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"lastSyncDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

