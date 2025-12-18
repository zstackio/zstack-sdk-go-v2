// Copyright (c) ZStack.io, Inc.

package param

// UpdateCdpTaskDetailParam UpdateCdpTask detail param
type UpdateCdpTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	BackupBandwidth int64 `json:"backupBandwidth,omitempty"`
	MaxCapacity int64 `json:"maxCapacity,omitempty"`
	MaxLatency int64 `json:"maxLatency,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateCdpTaskParam UpdateCdpTask request param
type UpdateCdpTaskParam struct {
	BaseParam
	Params UpdateCdpTaskDetailParam `json:"params"`
}
