// Copyright (c) ZStack.io, Inc.

package param

// SubmitLongJobDetailParam SubmitLongJob detail param
type SubmitLongJobDetailParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	JobName string `json:"jobName" validate:"required"`
	JobData string `json:"jobData" validate:"required"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SubmitLongJobParam SubmitLongJob request param
type SubmitLongJobParam struct {
	BaseParam
	Params SubmitLongJobDetailParam `json:"params"`
}
