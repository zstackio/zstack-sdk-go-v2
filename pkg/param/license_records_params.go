// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseRecordsDetailParam GetLicenseRecords详细参数
type GetLicenseRecordsDetailParam struct {
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
	rest bool `json:"count,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
}

// GetLicenseRecordsParam GetLicenseRecords请求参数
type GetLicenseRecordsParam struct {
	BaseParam
	Params GetLicenseRecordsDetailParam `json:"params"` // 详细参数
}

