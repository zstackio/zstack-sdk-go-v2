// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EmailTriggerActionInventoryView EmailTriggerAction
type EmailTriggerActionInventoryView struct {
	Email string `json:"email,omitempty"`
	MediaUuid string `json:"mediaUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Type string `json:"type,omitempty"`
	TriggerUuids []string `json:"triggerUuids,omitempty"`
}

