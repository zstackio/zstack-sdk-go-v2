// Copyright (c) ZStack.io, Inc.

package param

// CheckBatchDataIntegrityDetailParam CheckBatchDataIntegrity detail param
type CheckBatchDataIntegrityDetailParam struct {
	ResourceUuids []string `json:"resourceUuids,omitempty"`
	ResourceType string `json:"resourceType" validate:"required"`
}

// CheckBatchDataIntegrityParam CheckBatchDataIntegrity request param
type CheckBatchDataIntegrityParam struct {
	BaseParam
	Params CheckBatchDataIntegrityDetailParam `json:"params"`
}
