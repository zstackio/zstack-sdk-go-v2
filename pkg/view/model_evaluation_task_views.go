// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelEvaluationTaskInventoryView ModelEvaluationTask
type ModelEvaluationTaskInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"description,omitempty"`
	rest int `json:"percentage,omitempty"`
	rest string `json:"opaque,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"modelServiceGroupUuid,omitempty"`
	rest string `json:"evaluatedServiceGroupUuid,omitempty"`
	rest string `json:"datasetUuid,omitempty"`
	rest int `json:"limits,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int `json:"maxNewTokens,omitempty"`
	rest int `json:"topK,omitempty"`
	rest float32 `json:"temperature,omitempty"`
	rest float32 `json:"topP,omitempty"`
	rest string `json:"prompt,omitempty"`
	rest float32 `json:"repetitionPenalty,omitempty"`
	rest int `json:"maxLength,omitempty"`
	rest string `json:"model,omitempty"`
	rest string `json:"url,omitempty"`
	rest int `json:"parallel,omitempty"`
	rest int `json:"logEveryQuery,omitempty"`
	rest string `json:"api,omitempty"`
	rest map[string]string `json:"requestHeaders,omitempty"`
	rest int `json:"connectTimeout,omitempty"`
	rest int `json:"readTimeout,omitempty"`
}

