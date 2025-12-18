// Copyright (c) ZStack.io, Inc.

package param

// CreateSchedulerJobDetailParam CreateSchedulerJob detail param
type CreateSchedulerJobDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	Parameters map[string]string `json:"parameters,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSchedulerJobParam CreateSchedulerJob request param
type CreateSchedulerJobParam struct {
	BaseParam
	Params CreateSchedulerJobDetailParam `json:"params"`
}
