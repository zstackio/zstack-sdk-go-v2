// Copyright (c) ZStack.io, Inc.

package param

// AddV2VConversionHostDetailParam AddV2VConversionHost详细参数
type AddV2VConversionHostDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"storagePath" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddV2VConversionHostParam AddV2VConversionHost请求参数
type AddV2VConversionHostParam struct {
	BaseParam
	Params AddV2VConversionHostDetailParam `json:"params"` // 详细参数
}

