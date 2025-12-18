// Copyright (c) ZStack.io, Inc.

package param

// CreateMonitorTriggerDetailParam CreateMonitorTrigger detail param
type CreateMonitorTriggerDetailParam struct {
	Name string `json:"name" validate:"required"`
	Expression string `json:"expression" validate:"required"`
	Duration int `json:"duration" validate:"required"`
	RecoveryExpression string `json:"recoveryExpression,omitempty"`
	Description string `json:"description,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMonitorTriggerParam CreateMonitorTrigger request param
type CreateMonitorTriggerParam struct {
	BaseParam
	Params CreateMonitorTriggerDetailParam `json:"params"`
}
