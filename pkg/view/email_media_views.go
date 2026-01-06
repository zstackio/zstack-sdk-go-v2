// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EmailMediaInventoryView EmailMedia
type EmailMediaInventoryView struct {
	SmtpServer string `json:"smtpServer,omitempty"`
	SmtpPort int `json:"smtpPort,omitempty"`
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
}

// UpdateEmailMediaEventView UpdateEmailMediaEvent
type UpdateEmailMediaEventView struct {
	Inventory EmailMediaInventoryView `json:"inventory,omitempty"`
}

