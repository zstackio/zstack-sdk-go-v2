// Copyright (c) ZStack.io, Inc.

package view

// GetPrimaryStorageUsageReportView GetPrimaryStorageUsageReport
type GetPrimaryStorageUsageReportView struct {
	UriUsageForecast map[string]UsageReportView `json:"uriUsageForecast,omitempty"`
	UsageReport UsageReportView `json:"usageReport,omitempty"`
	Success bool `json:"success,omitempty"`
}

