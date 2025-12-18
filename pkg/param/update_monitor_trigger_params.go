// Copyright (c) ZStack.io, Inc.

package param

// UpdateMonitorTriggerDetailParam UpdateMonitorTrigger detail param
type UpdateMonitorTriggerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Expression string `json:"expression,omitempty"`
	Duration int `json:"duration,omitempty"`
}

// UpdateMonitorTriggerParam UpdateMonitorTrigger request param
type UpdateMonitorTriggerParam struct {
	BaseParam
	Params UpdateMonitorTriggerDetailParam `json:"params"`
}
