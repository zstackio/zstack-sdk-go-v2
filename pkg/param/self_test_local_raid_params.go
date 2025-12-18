// Copyright (c) ZStack.io, Inc.

package param

// SelfTestLocalRaidDetailParam SelfTestLocalRaid详细参数
type SelfTestLocalRaidDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SelfTestLocalRaidParam SelfTestLocalRaid请求参数
type SelfTestLocalRaidParam struct {
	BaseParam
	Params SelfTestLocalRaidDetailParam `json:"params"` // 详细参数
}

