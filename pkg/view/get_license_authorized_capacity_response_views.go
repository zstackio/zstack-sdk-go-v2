// Copyright (c) ZStack.io, Inc.

package view

// GetLicenseAuthorizedCapacityView GetLicenseAuthorizedCapacity
type GetLicenseAuthorizedCapacityView struct {
	Total interface{} `json:"total,omitempty"`
	Clients []interface{} `json:"clients,omitempty"`
	Server interface{} `json:"server,omitempty"`
	Success bool `json:"success,omitempty"`
}

