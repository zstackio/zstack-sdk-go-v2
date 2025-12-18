// Copyright (c) ZStack.io, Inc.

package param

// AddMonToCephPrimaryStorageDetailParam AddMonToCephPrimaryStorage详细参数
type AddMonToCephPrimaryStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"monUrls" validate:"required"` // 必填
}

// AddMonToCephPrimaryStorageParam AddMonToCephPrimaryStorage请求参数
type AddMonToCephPrimaryStorageParam struct {
	BaseParam
	Params AddMonToCephPrimaryStorageDetailParam `json:"params"` // 详细参数
}

