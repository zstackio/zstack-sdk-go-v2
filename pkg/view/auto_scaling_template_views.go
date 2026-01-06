// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingTemplateInventoryView AutoScalingTemplate
type AutoScalingTemplateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// DeleteAutoScalingTemplateEventView DeleteAutoScalingTemplateEvent
type DeleteAutoScalingTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

