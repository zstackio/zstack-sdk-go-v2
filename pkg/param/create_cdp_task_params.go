// Copyright (c) ZStack.io, Inc.

package param

// CreateCdpTaskDetailParam CreateCdpTask detail param
type CreateCdpTaskDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	TaskType string `json:"taskType" validate:"required"`
	PolicyUuid string `json:"policyUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	BackupBandwidth int64 `json:"backupBandwidth,omitempty"`
	MaxCapacity int64 `json:"maxCapacity,omitempty"`
	MaxLatency int64 `json:"maxLatency,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCdpTaskParam CreateCdpTask request param
type CreateCdpTaskParam struct {
	BaseParam
	Params CreateCdpTaskDetailParam `json:"params"`
}
