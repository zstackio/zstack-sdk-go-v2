// Copyright (c) ZStack.io, Inc.

package param

// StartBareMetal2InstanceDetailParam StartBareMetal2Instance detail param
type StartBareMetal2InstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	GatewayUuid string `json:"gatewayUuid,omitempty"`
	ChassisUuid string `json:"chassisUuid,omitempty"`
	ChassisOfferingUuid string `json:"chassisOfferingUuid,omitempty"`
}

// StartBareMetal2InstanceParam StartBareMetal2Instance request param
type StartBareMetal2InstanceParam struct {
	BaseParam
	Params StartBareMetal2InstanceDetailParam `json:"params"`
}
