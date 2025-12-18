// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2InstanceDetailParam UpdateBareMetal2Instance detail param
type UpdateBareMetal2InstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	ChassisOfferingUuid string `json:"chassisOfferingUuid,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	AutoReleaseChassisEvent string `json:"autoReleaseChassisEvent,omitempty"`
}

// UpdateBareMetal2InstanceParam UpdateBareMetal2Instance request param
type UpdateBareMetal2InstanceParam struct {
	BaseParam
	Params UpdateBareMetal2InstanceDetailParam `json:"params"`
}
