// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcHaGroupDetailParam CreateVpcHaGroup detail param
type CreateVpcHaGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	MonitorIps []string `json:"monitorIps,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcHaGroupParam CreateVpcHaGroup request param
type CreateVpcHaGroupParam struct {
	BaseParam
	Params CreateVpcHaGroupDetailParam `json:"params"`
}
