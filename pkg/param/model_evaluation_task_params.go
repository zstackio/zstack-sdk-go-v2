// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteModelEvaluationTaskParamDetail DeleteModelEvaluationTask detail param
type DeleteModelEvaluationTaskParamDetail struct {
}

// DeleteModelEvaluationTaskParam DeleteModelEvaluationTask request param
type DeleteModelEvaluationTaskParam struct {
	BaseParam
	Params DeleteModelEvaluationTaskParamDetail `json:"deleteModelEvaluationTask"`
}
// UpdateModelEvaluationTaskParamDetail UpdateModelEvaluationTask detail param
type UpdateModelEvaluationTaskParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateModelEvaluationTaskParam UpdateModelEvaluationTask request param
type UpdateModelEvaluationTaskParam struct {
	BaseParam
	Params UpdateModelEvaluationTaskParamDetail `json:"updateModelEvaluationTask"`
}
