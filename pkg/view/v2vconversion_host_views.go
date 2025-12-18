// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// V2VConversionHostInventoryView V2VConversionHost
type V2VConversionHostInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	StoragePath string `json:"storagePath,omitempty"`
	State string `json:"state,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
	AvailableSize int64 `json:"availableSize,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	HostStatus string `json:"hostStatus,omitempty"`
	HostState string `json:"hostState,omitempty"`
}

