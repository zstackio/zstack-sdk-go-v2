// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSDingTalkEndpointInventoryView SNSDingTalkEndpoint
type SNSDingTalkEndpointInventoryView struct {
	rest string `json:"url,omitempty"`
	rest bool `json:"atAll,omitempty"`
	rest string `json:"secret,omitempty"`
	rest []string `json:"atPersonPhoneNumbers,omitempty"`
	rest []SNSDingTalkAtPersonInventoryView `json:"atPersonList,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"platformUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"connectionStatus,omitempty"`
	rest SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

