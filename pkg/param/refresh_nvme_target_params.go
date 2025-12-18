// Copyright (c) ZStack.io, Inc.

package param

// RefreshNvmeTargetDetailParam RefreshNvmeTarget detail param
type RefreshNvmeTargetDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	NvmeLunUuids []string `json:"nvmeLunUuids,omitempty"`
}

// RefreshNvmeTargetParam RefreshNvmeTarget request param
type RefreshNvmeTargetParam struct {
	BaseParam
	Params RefreshNvmeTargetDetailParam `json:"params"`
}
