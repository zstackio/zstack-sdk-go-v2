// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelEvaluationTaskDetailParam UpdateModelEvaluationTask detail param
type UpdateModelEvaluationTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateModelEvaluationTaskParam UpdateModelEvaluationTask request param
type UpdateModelEvaluationTaskParam struct {
	BaseParam
	Params UpdateModelEvaluationTaskDetailParam `json:"params"`
}
