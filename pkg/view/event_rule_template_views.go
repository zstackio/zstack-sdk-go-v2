// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EventRuleTemplateInventoryView EventRuleTemplate
type EventRuleTemplateInventoryView struct {
	Name string `json:"name,omitempty"`
	MonitorTemplateUuid string `json:"monitorTemplateUuid,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	EventName string `json:"eventName,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels string `json:"labels,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

