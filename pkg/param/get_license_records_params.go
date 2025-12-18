// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseRecordsDetailParam GetLicenseRecords detail param
type GetLicenseRecordsDetailParam struct {
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
	ReplyWithCount bool `json:"replyWithCount,omitempty"`
	Count bool `json:"count,omitempty"`
	SortBy string `json:"sortBy,omitempty"`
	SortDirection string `json:"sortDirection,omitempty"`
}

// GetLicenseRecordsParam GetLicenseRecords request param
type GetLicenseRecordsParam struct {
	BaseParam
	Params GetLicenseRecordsDetailParam `json:"params"`
}
