// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateSchedulerJobGroupParamDetail CreateSchedulerJobGroup detail param
type CreateSchedulerJobGroupParamDetail struct {
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
	CreateSchedulerJobGroup CreateSchedulerJobGroupParamDetail `json:"createSchedulerJobGroup"`
}
// DeleteSchedulerJobGroupParamDetail DeleteSchedulerJobGroup detail param
type DeleteSchedulerJobGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSchedulerJobGroupParam DeleteSchedulerJobGroup request param
type DeleteSchedulerJobGroupParam struct {
	BaseParam
	DeleteSchedulerJobGroup DeleteSchedulerJobGroupParamDetail `json:"deleteSchedulerJobGroup"`
}
// UpdateSchedulerJobGroupParamDetail UpdateSchedulerJobGroup detail param
type UpdateSchedulerJobGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// UpdateSchedulerJobGroupParam UpdateSchedulerJobGroup request param
type UpdateSchedulerJobGroupParam struct {
	BaseParam
	UpdateSchedulerJobGroup UpdateSchedulerJobGroupParamDetail `json:"updateSchedulerJobGroup"`
}
