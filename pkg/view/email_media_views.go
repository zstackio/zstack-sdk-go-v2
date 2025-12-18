// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EmailMediaInventoryView EmailMedia
type EmailMediaInventoryView struct {
	rest string `json:"smtpServer,omitempty"`
	rest int `json:"smtpPort,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

