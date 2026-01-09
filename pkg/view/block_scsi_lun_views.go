// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BlockScsiLunInventoryView BlockScsiLun
type BlockScsiLunInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Wwn *string `json:"wwn,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Id *int `json:"id,omitempty"`
	VolumeUuid *string `json:"volumeUuid,omitempty"`
	Target *string `json:"target,omitempty"`
	LunMapId *int `json:"lunMapId,omitempty"`
	LunType *string `json:"lunType,omitempty"`
	LunInitSnapshotID *int `json:"lunInitSnapshotID,omitempty"`
	UsedSize int64 `json:"usedSize,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

