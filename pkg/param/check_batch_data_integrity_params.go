// Copyright (c) ZStack.io, Inc.

package param

// CheckBatchDataIntegrityDetailParam CheckBatchDataIntegrity详细参数
type CheckBatchDataIntegrityDetailParam struct {
	rest []string `json:"resourceUuids,omitempty"`
	rest string `json:"resourceType" validate:"required"` // 必填
}

// CheckBatchDataIntegrityParam CheckBatchDataIntegrity请求参数
type CheckBatchDataIntegrityParam struct {
	BaseParam
	Params CheckBatchDataIntegrityDetailParam `json:"params"` // 详细参数
}

