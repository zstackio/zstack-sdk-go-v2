// Copyright (c) ZStack.io, Inc.

package param

// RecoverResourceSplitBrainDetailParam RecoverResourceSplitBrain详细参数
type RecoverResourceSplitBrainDetailParam struct {
	rest string `json:"resourceUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"primaryStorageUuid,omitempty"`
	rest bool `json:"forceRecover,omitempty"`
	rest string `json:"resourceType,omitempty"`
}

// RecoverResourceSplitBrainParam RecoverResourceSplitBrain请求参数
type RecoverResourceSplitBrainParam struct {
	BaseParam
	Params RecoverResourceSplitBrainDetailParam `json:"params"` // 详细参数
}

