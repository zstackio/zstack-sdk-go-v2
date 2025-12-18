// Copyright (c) ZStack.io, Inc.

package param

// AddEventRuleTemplateDetailParam AddEventRuleTemplate详细参数
type AddEventRuleTemplateDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"monitorTemplateUuid" validate:"required"` // 必填
	rest string `json:"namespace" validate:"required"` // 必填
	rest string `json:"eventName" validate:"required"` // 必填
	rest string `json:"emergencyLevel,omitempty"`
	rest []interface{} `json:"labels,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddEventRuleTemplateParam AddEventRuleTemplate请求参数
type AddEventRuleTemplateParam struct {
	BaseParam
	Params AddEventRuleTemplateDetailParam `json:"params"` // 详细参数
}

