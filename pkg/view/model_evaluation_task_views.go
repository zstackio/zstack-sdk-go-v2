// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelEvaluationTaskInventoryView ModelEvaluationTask
type ModelEvaluationTaskInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Percentage int `json:"percentage,omitempty"`
	Opaque string `json:"opaque,omitempty"`
	Status string `json:"status,omitempty"`
	ModelServiceGroupUuid string `json:"modelServiceGroupUuid,omitempty"`
	EvaluatedServiceGroupUuid string `json:"evaluatedServiceGroupUuid,omitempty"`
	DatasetUuid string `json:"datasetUuid,omitempty"`
	Limits int `json:"limits,omitempty"`
	MaxNewTokens int `json:"maxNewTokens,omitempty"`
	TopK int `json:"topK,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	TopP float32 `json:"topP,omitempty"`
	Prompt string `json:"prompt,omitempty"`
	RepetitionPenalty float32 `json:"repetitionPenalty,omitempty"`
	MaxLength int `json:"maxLength,omitempty"`
	Model string `json:"model,omitempty"`
	Url string `json:"url,omitempty"`
	Parallel int `json:"parallel,omitempty"`
	LogEveryQuery int `json:"logEveryQuery,omitempty"`
	Api string `json:"api,omitempty"`
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`
	ConnectTimeout int `json:"connectTimeout,omitempty"`
	ReadTimeout int `json:"readTimeout,omitempty"`
	TotalScore float64 `json:"totalScore,omitempty"`
	EndTime time.Time `json:"endTime,omitempty"`
}

// DeleteModelEvaluationTaskEventView DeleteModelEvaluationTaskEvent
type DeleteModelEvaluationTaskEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryModelEvaluationTaskView QueryModelEvaluationTask
type QueryModelEvaluationTaskView struct {
	Inventories []ModelEvaluationTaskInventoryView `json:"inventories,omitempty"`
}

// UpdateModelEvaluationTaskEventView UpdateModelEvaluationTaskEvent
type UpdateModelEvaluationTaskEventView struct {
	Inventory ModelEvaluationTaskInventoryView `json:"inventory,omitempty"`
}

