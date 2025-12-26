// Copyright (c) ZStack.io, Inc.

package param

// UpdateEventRuleTemplateDetailParam UpdateEventRuleTemplate detail param
type UpdateEventRuleTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
}

// UpdateEventRuleTemplateParam UpdateEventRuleTemplate request param
type UpdateEventRuleTemplateParam struct {
	BaseParam
	Params UpdateEventRuleTemplateDetailParam `json:"params"`
}
