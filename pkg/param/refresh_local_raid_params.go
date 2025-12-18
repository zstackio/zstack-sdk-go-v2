// Copyright (c) ZStack.io, Inc.

package param

// RefreshLocalRaidDetailParam RefreshLocalRaid详细参数
type RefreshLocalRaidDetailParam struct {
	rest string `json:"hostUuid" validate:"required"` // 必填
}

// RefreshLocalRaidParam RefreshLocalRaid请求参数
type RefreshLocalRaidParam struct {
	BaseParam
	Params RefreshLocalRaidDetailParam `json:"params"` // 详细参数
}

