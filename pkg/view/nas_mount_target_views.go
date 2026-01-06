// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NasMountTargetInventoryView NasMountTarget
type NasMountTargetInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MountDomain string `json:"mountDomain,omitempty"`
	NasFileSystemUuid string `json:"nasFileSystemUuid,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// UpdateNasMountTargetEventView UpdateNasMountTargetEvent
type UpdateNasMountTargetEventView struct {
	Inventory NasMountTargetInventoryView `json:"inventory,omitempty"`
}

// QueryNasMountTargetView QueryNasMountTarget
type QueryNasMountTargetView struct {
	Inventories []NasMountTargetInventoryView `json:"inventories,omitempty"`
}

// DeleteNasMountTargetEventView DeleteNasMountTargetEvent
type DeleteNasMountTargetEventView struct {
	Success bool `json:"success,omitempty"`
}

