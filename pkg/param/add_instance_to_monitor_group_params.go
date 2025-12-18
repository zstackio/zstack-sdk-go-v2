// Copyright (c) ZStack.io, Inc.

package param

// AddInstanceToMonitorGroupDetailParam AddInstanceToMonitorGroup detail param
type AddInstanceToMonitorGroupDetailParam struct {
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddInstanceToMonitorGroupParam AddInstanceToMonitorGroup request param
type AddInstanceToMonitorGroupParam struct {
	BaseParam
	Params AddInstanceToMonitorGroupDetailParam `json:"params"`
}
