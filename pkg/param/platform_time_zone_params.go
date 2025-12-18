// Copyright (c) ZStack.io, Inc.

package param

// GetPlatformTimeZoneDetailParam GetPlatformTimeZone详细参数
type GetPlatformTimeZoneDetailParam struct {
}

// GetPlatformTimeZoneParam GetPlatformTimeZone请求参数
type GetPlatformTimeZoneParam struct {
	BaseParam
	Params GetPlatformTimeZoneDetailParam `json:"params"` // 详细参数
}

