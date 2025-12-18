// Copyright (c) ZStack.io, Inc.

package param

// CreateFlowCollectorDetailParam CreateFlowCollector detail param
type CreateFlowCollectorDetailParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FlowMeterUuid string `json:"flowMeterUuid" validate:"required"`
	Server string `json:"server,omitempty"`
	Port int64 `json:"port,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFlowCollectorParam CreateFlowCollector request param
type CreateFlowCollectorParam struct {
	BaseParam
	Params CreateFlowCollectorDetailParam `json:"params"`
}
