// Copyright (c) ZStack.io, Inc.

package param

// GenerateModelMetadataDetailParam GenerateModelMetadata detail param
type GenerateModelMetadataDetailParam struct {
	ModelCenterUuid string `json:"modelCenterUuid" validate:"required"`
	ModelUuids []string `json:"modelUuids,omitempty"`
}

// GenerateModelMetadataParam GenerateModelMetadata request param
type GenerateModelMetadataParam struct {
	BaseParam
	Params GenerateModelMetadataDetailParam `json:"params"`
}
