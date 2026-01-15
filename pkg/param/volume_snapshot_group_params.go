// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateVolumeSnapshotGroupParamDetail CreateVolumeSnapshotGroup detail param
type CreateVolumeSnapshotGroupParamDetail struct {
	RootVolumeUuid string `json:"rootVolumeUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	WithMemory bool `json:"withMemory,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVolumeSnapshotGroupParam CreateVolumeSnapshotGroup request param
type CreateVolumeSnapshotGroupParam struct {
	BaseParam
	CreateVolumeSnapshotGroup CreateVolumeSnapshotGroupParamDetail `json:"createVolumeSnapshotGroup"`
}
// UpdateVolumeSnapshotGroupParamDetail UpdateVolumeSnapshotGroup detail param
type UpdateVolumeSnapshotGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
}

// UpdateVolumeSnapshotGroupParam UpdateVolumeSnapshotGroup request param
type UpdateVolumeSnapshotGroupParam struct {
	BaseParam
	UpdateVolumeSnapshotGroup UpdateVolumeSnapshotGroupParamDetail `json:"updateVolumeSnapshotGroup"`
}
// DeleteVolumeSnapshotGroupParamDetail DeleteVolumeSnapshotGroup detail param
type DeleteVolumeSnapshotGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Direction string `json:"direction,omitempty"`
	Scope string `json:"scope,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVolumeSnapshotGroupParam DeleteVolumeSnapshotGroup request param
type DeleteVolumeSnapshotGroupParam struct {
	BaseParam
	DeleteVolumeSnapshotGroup DeleteVolumeSnapshotGroupParamDetail `json:"deleteVolumeSnapshotGroup"`
}
