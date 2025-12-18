// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelEvaluationTasksDetailParam DeleteModelEvaluationTasks详细参数
type DeleteModelEvaluationTasksDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
}

// DeleteModelEvaluationTasksParam DeleteModelEvaluationTasks请求参数
type DeleteModelEvaluationTasksParam struct {
	BaseParam
	Params DeleteModelEvaluationTasksDetailParam `json:"params"` // 详细参数
}

