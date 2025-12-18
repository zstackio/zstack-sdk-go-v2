// Copyright (c) ZStack.io, Inc.

package param

// CalculateImageHashDetailParam CalculateImageHash详细参数
type CalculateImageHashDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"algorithm,omitempty"`
}

// CalculateImageHashParam CalculateImageHash请求参数
type CalculateImageHashParam struct {
	BaseParam
	Params CalculateImageHashDetailParam `json:"params"` // 详细参数
}

