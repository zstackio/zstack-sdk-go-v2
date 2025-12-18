// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelEvaluationTaskDetailParam DeleteModelEvaluationTask detail param
type DeleteModelEvaluationTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteModelEvaluationTaskParam DeleteModelEvaluationTask request param
type DeleteModelEvaluationTaskParam struct {
	BaseParam
	Params DeleteModelEvaluationTaskDetailParam `json:"params"`
}
