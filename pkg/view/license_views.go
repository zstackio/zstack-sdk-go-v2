// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseInventoryView License
type LicenseInventoryView struct {
	Uuid               string               `json:"uuid,omitempty"`
	User               string               `json:"user,omitempty"`
	ProdInfo           string               `json:"prodInfo,omitempty"`
	CpuNum             int                  `json:"cpuNum,omitempty"`
	HostNum            int                  `json:"hostNum,omitempty"`
	VmNum              int                  `json:"vmNum,omitempty"`
	Capacity           int                  `json:"capacity,omitempty"`
	LicenseType        string               `json:"licenseType,omitempty"`
	QuotaType          string               `json:"quotaType,omitempty"`
	ExpiredDate        string               `json:"expiredDate,omitempty"`
	IssuedDate         string               `json:"issuedDate,omitempty"`
	UploadDate         string               `json:"uploadDate,omitempty"`
	ManagementNodeUuid string               `json:"managementNodeUuid,omitempty"`
	Expired            bool                 `json:"expired,omitempty"`
	Source             string               `json:"source,omitempty"`
	PlatformId         string               `json:"platformId,omitempty"`
	LicenseRequest     string               `json:"licenseRequest,omitempty"`
	AvailableHostNum   int                  `json:"availableHostNum,omitempty"`
	AvailableCpuNum    int                  `json:"availableCpuNum,omitempty"`
	AvailableVmNum     int                  `json:"availableVmNum,omitempty"`
	Usage              LicenseUsageViewView `json:"usage,omitempty"`
}

// GetLicenseRecordsView GetLicenseRecords
type GetLicenseRecordsView struct {
	Inventories []LicenseInventoryView `json:"inventories,omitempty"`
	Total       int64                  `json:"total,omitempty"`
	Success     bool                   `json:"success,omitempty"`
}

// DeleteLicenseEventView DeleteLicenseEvent
type DeleteLicenseEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateLicenseEventView UpdateLicenseEvent
type UpdateLicenseEventView struct {
	Inventory LicenseInventoryView `json:"inventory,omitempty"`
}

// GetLicenseInfoView GetLicenseInfo
type GetLicenseInfoView struct {
	Inventory LicenseInventoryView        `json:"inventory,omitempty"`
	Additions []AdditionalLicenseInfoView `json:"additions,omitempty"`
}
