// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasMountTargetInventoryView AliyunNasMountTarget
type AliyunNasMountTargetInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccessGroupUuid *string `json:"accessGroupUuid,omitempty"`
	Status *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
	MountDomain *string `json:"mountDomain,omitempty"`
	NasFileSystemUuid *string `json:"nasFileSystemUuid,omitempty"`
	Type *string `json:"type,omitempty"`
}

// AddAliyunNasMountTargetEventView AddAliyunNasMountTargetEvent
type AddAliyunNasMountTargetEventView struct {
	Inventory AliyunNasMountTargetInventoryView `json:"inventory,omitempty"`
}

