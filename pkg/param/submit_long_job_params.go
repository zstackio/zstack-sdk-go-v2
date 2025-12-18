// Copyright (c) ZStack.io, Inc.

package param

// SubmitLongJobDetailParam SubmitLongJob详细参数
type SubmitLongJobDetailParam struct {
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"jobName" validate:"required"` // 必填
	rest string `json:"jobData" validate:"required"` // 必填
	rest string `json:"targetResourceUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SubmitLongJobParam SubmitLongJob请求参数
type SubmitLongJobParam struct {
	BaseParam
	Params SubmitLongJobDetailParam `json:"params"` // 详细参数
}

