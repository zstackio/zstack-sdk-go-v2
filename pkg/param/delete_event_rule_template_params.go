// Copyright (c) ZStack.io, Inc.

package param

// DeleteEventRuleTemplateDetailParam DeleteEventRuleTemplate detail param
type DeleteEventRuleTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEventRuleTemplateParam DeleteEventRuleTemplate request param
type DeleteEventRuleTemplateParam struct {
	BaseParam
	Params DeleteEventRuleTemplateDetailParam `json:"params"`
}
