// Copyright (c) ZStack.io, Inc.

package param

// CreateBlockVolumeDetailParam CreateBlockVolume detail param
type CreateBlockVolumeDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Size int64 `json:"size" validate:"required"`
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	AccessPathId int `json:"accessPathId,omitempty"`
	AccessPathIqn string `json:"accessPathIqn,omitempty"`
	BurstTotalBw int64 `json:"burstTotalBw,omitempty"`
	BurstTotalIops int64 `json:"burstTotalIops,omitempty"`
	MaxTotalBw int64 `json:"maxTotalBw,omitempty"`
	MaxTotalIops int64 `json:"maxTotalIops,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBlockVolumeParam CreateBlockVolume request param
type CreateBlockVolumeParam struct {
	BaseParam
	Params CreateBlockVolumeDetailParam `json:"params"`
}
