// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddEventRuleTemplateParamDetail AddEventRuleTemplate detail param
type AddEventRuleTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	MonitorTemplateUuid string `json:"monitorTemplateUuid" validate:"required"`
	Namespace string `json:"namespace" validate:"required"`
	EventName string `json:"eventName" validate:"required"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddEventRuleTemplateParam AddEventRuleTemplate request param
type AddEventRuleTemplateParam struct {
	BaseParam
	AddEventRuleTemplate AddEventRuleTemplateParamDetail `json:"addEventRuleTemplate"`
}
// DeleteEventRuleTemplateParamDetail DeleteEventRuleTemplate detail param
type DeleteEventRuleTemplateParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEventRuleTemplateParam DeleteEventRuleTemplate request param
type DeleteEventRuleTemplateParam struct {
	BaseParam
	DeleteEventRuleTemplate DeleteEventRuleTemplateParamDetail `json:"deleteEventRuleTemplate"`
}
// UpdateEventRuleTemplateParamDetail UpdateEventRuleTemplate detail param
type UpdateEventRuleTemplateParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
}

// UpdateEventRuleTemplateParam UpdateEventRuleTemplate request param
type UpdateEventRuleTemplateParam struct {
	BaseParam
	UpdateEventRuleTemplate UpdateEventRuleTemplateParamDetail `json:"updateEventRuleTemplate"`
}
