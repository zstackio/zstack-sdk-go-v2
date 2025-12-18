// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSEmailPlatformInventoryView SNSEmailPlatform
type SNSEmailPlatformInventoryView struct {
	rest string `json:"smtpServer,omitempty"`
	rest int `json:"smtpPort,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

