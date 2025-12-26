// Copyright (c) ZStack.io, Inc.

package param

// SubscribeEventDetailParam SubscribeEvent detail param
type SubscribeEventDetailParam struct {
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace" validate:"required"`
	EventName string `json:"eventName" validate:"required"`
	Actions []ActionParamParam `json:"actions,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SubscribeEventParam SubscribeEvent request param
type SubscribeEventParam struct {
	BaseParam
	Params SubscribeEventDetailParam `json:"params"`
}
