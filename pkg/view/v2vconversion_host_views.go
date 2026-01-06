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
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	HostStatus string `json:"hostStatus,omitempty"`
	HostState string `json:"hostState,omitempty"`
}

// AddV2VConversionHostEventView AddV2VConversionHostEvent
type AddV2VConversionHostEventView struct {
	Inventory V2VConversionHostInventoryView `json:"inventory,omitempty"`
}

// UpdateV2VConversionHostEventView UpdateV2VConversionHostEvent
type UpdateV2VConversionHostEventView struct {
	Inventory V2VConversionHostInventoryView `json:"inventory,omitempty"`
}

// QueryV2VConversionHostView QueryV2VConversionHost
type QueryV2VConversionHostView struct {
	Inventories []V2VConversionHostInventoryView `json:"inventories,omitempty"`
}

// DeleteV2VConversionHostEventView DeleteV2VConversionHostEvent
type DeleteV2VConversionHostEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeV2VConversionHostStateEventView ChangeV2VConversionHostStateEvent
type ChangeV2VConversionHostStateEventView struct {
	Inventory V2VConversionHostInventoryView `json:"inventory,omitempty"`
}

