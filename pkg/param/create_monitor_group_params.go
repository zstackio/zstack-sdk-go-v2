// Copyright (c) ZStack.io, Inc.

package param

// CreateMonitorGroupDetailParam CreateMonitorGroup detail param
type CreateMonitorGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Actions []ActionParamParam `json:"actions,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMonitorGroupParam CreateMonitorGroup request param
type CreateMonitorGroupParam struct {
	BaseParam
	Params CreateMonitorGroupDetailParam `json:"params"`
}
