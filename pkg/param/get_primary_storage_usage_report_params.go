// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageUsageReportDetailParam GetPrimaryStorageUsageReport detail param
type GetPrimaryStorageUsageReportDetailParam struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	Uris []string `json:"uris,omitempty"`
}

// GetPrimaryStorageUsageReportParam GetPrimaryStorageUsageReport request param
type GetPrimaryStorageUsageReportParam struct {
	BaseParam
	Params GetPrimaryStorageUsageReportDetailParam `json:"params"`
}
