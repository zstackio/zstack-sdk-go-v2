// Copyright (c) ZStack.io, Inc.

package param

// RefreshLocalRaidDetailParam RefreshLocalRaid detail param
type RefreshLocalRaidDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
}

// RefreshLocalRaidParam RefreshLocalRaid request param
type RefreshLocalRaidParam struct {
	BaseParam
	Params RefreshLocalRaidDetailParam `json:"params"`
}
