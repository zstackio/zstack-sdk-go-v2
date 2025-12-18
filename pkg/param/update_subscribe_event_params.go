// Copyright (c) ZStack.io, Inc.

package param

// UpdateSubscribeEventDetailParam UpdateSubscribeEvent detail param
type UpdateSubscribeEventDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateSubscribeEventParam UpdateSubscribeEvent request param
type UpdateSubscribeEventParam struct {
	BaseParam
	Params UpdateSubscribeEventDetailParam `json:"params"`
}
