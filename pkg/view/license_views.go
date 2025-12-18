// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LicenseInventoryView License
type LicenseInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"user,omitempty"`
	rest string `json:"prodInfo,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int `json:"hostNum,omitempty"`
	rest int `json:"vmNum,omitempty"`
	rest int `json:"capacity,omitempty"`
	rest string `json:"licenseType,omitempty"`
	rest string `json:"quotaType,omitempty"`
	rest string `json:"expiredDate,omitempty"`
	rest string `json:"issuedDate,omitempty"`
	rest string `json:"uploadDate,omitempty"`
	rest string `json:"managementNodeUuid,omitempty"`
	rest bool `json:"expired,omitempty"`
	rest string `json:"source,omitempty"`
	rest string `json:"platformId,omitempty"`
	rest string `json:"licenseRequest,omitempty"`
	rest int `json:"availableHostNum,omitempty"`
	rest int `json:"availableCpuNum,omitempty"`
	rest int `json:"availableVmNum,omitempty"`
	rest interface{} `json:"usage,omitempty"`
}

