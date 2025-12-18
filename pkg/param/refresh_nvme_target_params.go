// Copyright (c) ZStack.io, Inc.

package param

// RefreshNvmeTargetDetailParam RefreshNvmeTarget详细参数
type RefreshNvmeTargetDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest []string `json:"nvmeLunUuids,omitempty"`
}

// RefreshNvmeTargetParam RefreshNvmeTarget请求参数
type RefreshNvmeTargetParam struct {
	BaseParam
	Params RefreshNvmeTargetDetailParam `json:"params"` // 详细参数
}

