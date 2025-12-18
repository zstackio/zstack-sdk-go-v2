// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ThirdpartyPlatformInventoryView ThirdpartyPlatform
type ThirdpartyPlatformInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Template string `json:"template,omitempty"`
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	LastSyncDate time.Time `json:"lastSyncDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

