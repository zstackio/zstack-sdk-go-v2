// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSUniversalSmsEndpointInventoryView SNSUniversalSmsEndpoint
type SNSUniversalSmsEndpointInventoryView struct {
	rest string `json:"smsAccessKeyId,omitempty"`
	rest string `json:"smsAccessKeySecret,omitempty"`
	rest string `json:"supplier,omitempty"`
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

