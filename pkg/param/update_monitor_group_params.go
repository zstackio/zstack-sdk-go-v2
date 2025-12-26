// Copyright (c) ZStack.io, Inc.

package param

// UpdateMonitorGroupDetailParam UpdateMonitorGroup detail param
type UpdateMonitorGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Actions []ActionParamParam `json:"actions,omitempty"`
	StateEvent string `json:"stateEvent,omitempty"`
}

// UpdateMonitorGroupParam UpdateMonitorGroup request param
type UpdateMonitorGroupParam struct {
	BaseParam
	Params UpdateMonitorGroupDetailParam `json:"params"`
}
