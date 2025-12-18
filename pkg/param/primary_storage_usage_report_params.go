// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageUsageReportDetailParam GetPrimaryStorageUsageReport详细参数
type GetPrimaryStorageUsageReportDetailParam struct {
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
	rest []string `json:"uris,omitempty"`
}

// GetPrimaryStorageUsageReportParam GetPrimaryStorageUsageReport请求参数
type GetPrimaryStorageUsageReportParam struct {
	BaseParam
	Params GetPrimaryStorageUsageReportDetailParam `json:"params"` // 详细参数
}

