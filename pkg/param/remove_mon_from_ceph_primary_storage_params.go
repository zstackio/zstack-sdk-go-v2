// Copyright (c) ZStack.io, Inc.

package param

// RemoveMonFromCephPrimaryStorageDetailParam RemoveMonFromCephPrimaryStorage详细参数
type RemoveMonFromCephPrimaryStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"monHostnames" validate:"required"` // 必填
}

// RemoveMonFromCephPrimaryStorageParam RemoveMonFromCephPrimaryStorage请求参数
type RemoveMonFromCephPrimaryStorageParam struct {
	BaseParam
	Params RemoveMonFromCephPrimaryStorageDetailParam `json:"params"` // 详细参数
}

