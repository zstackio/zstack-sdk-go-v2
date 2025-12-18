// Copyright (c) ZStack.io, Inc.

package view

// VerifyLicenseServerEventView VerifyLicenseServerEvent
type VerifyLicenseServerEventView struct {
	AccessKeyId string `json:"accessKeyId,omitempty"`
	AccessKeySecret string `json:"accessKeySecret,omitempty"`
	LicenseClient LicenseAuthorizedNodeInventoryView `json:"licenseClient,omitempty"`
	LicenseServer LicenseAuthorizedNodeInventoryView `json:"licenseServer,omitempty"`
	Success bool `json:"success,omitempty"`
}

