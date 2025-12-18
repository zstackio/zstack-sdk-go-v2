// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelEvaluationTasksDetailParam DeleteModelEvaluationTasks detail param
type DeleteModelEvaluationTasksDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// DeleteModelEvaluationTasksParam DeleteModelEvaluationTasks request param
type DeleteModelEvaluationTasksParam struct {
	BaseParam
	Params DeleteModelEvaluationTasksDetailParam `json:"params"`
}
