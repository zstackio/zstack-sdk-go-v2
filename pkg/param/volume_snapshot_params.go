// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateVolumeSnapshotParamDetail UpdateVolumeSnapshot detail param
type UpdateVolumeSnapshotParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVolumeSnapshotParam UpdateVolumeSnapshot request param
type UpdateVolumeSnapshotParam struct {
	BaseParam
	UpdateVolumeSnapshot UpdateVolumeSnapshotParamDetail `json:"updateVolumeSnapshot"`
}
// DeleteVolumeSnapshotParamDetail DeleteVolumeSnapshot detail param
type DeleteVolumeSnapshotParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Direction string `json:"direction,omitempty"`
	Scope string `json:"scope,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVolumeSnapshotParam DeleteVolumeSnapshot request param
type DeleteVolumeSnapshotParam struct {
	BaseParam
	DeleteVolumeSnapshot DeleteVolumeSnapshotParamDetail `json:"deleteVolumeSnapshot"`
}
// CreateVolumeSnapshotParamDetail CreateVolumeSnapshot detail param
type CreateVolumeSnapshotParamDetail struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVolumeSnapshotParam CreateVolumeSnapshot request param
type CreateVolumeSnapshotParam struct {
	BaseParam
	CreateVolumeSnapshot CreateVolumeSnapshotParamDetail `json:"createVolumeSnapshot"`
}
