// Copyright (c) ZStack.io, Inc.

package param

// CreateSchedulerJobGroupDetailParam CreateSchedulerJobGroup detail param
type CreateSchedulerJobGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	Parameters map[string]string `json:"parameters,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSchedulerJobGroupParam CreateSchedulerJobGroup request param
type CreateSchedulerJobGroupParam struct {
	BaseParam
	Params CreateSchedulerJobGroupDetailParam `json:"params"`
}
