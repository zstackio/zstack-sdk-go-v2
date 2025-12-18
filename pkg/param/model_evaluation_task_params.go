// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelEvaluationTaskDetailParam UpdateModelEvaluationTask详细参数
type UpdateModelEvaluationTaskDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateModelEvaluationTaskParam UpdateModelEvaluationTask请求参数
type UpdateModelEvaluationTaskParam struct {
	BaseParam
	Params UpdateModelEvaluationTaskDetailParam `json:"params"` // 详细参数
}

