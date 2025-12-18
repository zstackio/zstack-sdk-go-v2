// Copyright (c) ZStack.io, Inc.

package param

// AddEventRuleTemplateDetailParam AddEventRuleTemplate detail param
type AddEventRuleTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	MonitorTemplateUuid string `json:"monitorTemplateUuid" validate:"required"`
	Namespace string `json:"namespace" validate:"required"`
	EventName string `json:"eventName" validate:"required"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels []interface{} `json:"labels,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddEventRuleTemplateParam AddEventRuleTemplate request param
type AddEventRuleTemplateParam struct {
	BaseParam
	Params AddEventRuleTemplateDetailParam `json:"params"`
}
