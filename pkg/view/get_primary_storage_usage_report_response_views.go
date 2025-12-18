// Copyright (c) ZStack.io, Inc.

package view

// GetPrimaryStorageUsageReportView GetPrimaryStorageUsageReport
type GetPrimaryStorageUsageReportView struct {
	UriUsageForecast map[string]interface{} `json:"uriUsageForecast,omitempty"`
	UsageReport interface{} `json:"usageReport,omitempty"`
	Success bool `json:"success,omitempty"`
}

