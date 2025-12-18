// Copyright (c) ZStack.io, Inc.

package param

// SelfTestLocalRaidDetailParam SelfTestLocalRaid detail param
type SelfTestLocalRaidDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SelfTestLocalRaidParam SelfTestLocalRaid request param
type SelfTestLocalRaidParam struct {
	BaseParam
	Params SelfTestLocalRaidDetailParam `json:"params"`
}
