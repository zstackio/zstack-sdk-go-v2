// Copyright (c) ZStack.io, Inc.

package view

// RegisterLicenseRequestedApplicationEventView RegisterLicenseRequestedApplicationEvent
type RegisterLicenseRequestedApplicationEventView struct {
	AppId string `json:"appId,omitempty"`
	ServicePubKey string `json:"servicePubKey,omitempty"`
	Success bool `json:"success,omitempty"`
}

