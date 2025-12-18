// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasMountTargetInventoryView AliyunNasMountTarget
type AliyunNasMountTargetInventoryView struct {
	AccessGroupUuid string `json:"accessGroupUuid,omitempty"`
	Status string `json:"status,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MountDomain string `json:"mountDomain,omitempty"`
	NasFileSystemUuid string `json:"nasFileSystemUuid,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

