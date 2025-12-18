// Copyright (c) ZStack.io, Inc.

package param

// GenerateModelMetadataDetailParam GenerateModelMetadata详细参数
type GenerateModelMetadataDetailParam struct {
	rest string `json:"modelCenterUuid" validate:"required"` // 必填
	rest []string `json:"modelUuids,omitempty"`
}

// GenerateModelMetadataParam GenerateModelMetadata请求参数
type GenerateModelMetadataParam struct {
	BaseParam
	Params GenerateModelMetadataDetailParam `json:"params"` // 详细参数
}

