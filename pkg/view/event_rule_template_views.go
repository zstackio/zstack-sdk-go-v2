// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EventRuleTemplateInventoryView EventRuleTemplate
type EventRuleTemplateInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"monitorTemplateUuid,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"eventName,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"labels,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"uuid,omitempty"`
}

