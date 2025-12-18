// Copyright (c) ZStack.io, Inc.

package param

// GetPlatformTimeZoneDetailParam GetPlatformTimeZone detail param
type GetPlatformTimeZoneDetailParam struct {
}

// GetPlatformTimeZoneParam GetPlatformTimeZone request param
type GetPlatformTimeZoneParam struct {
	BaseParam
	Params GetPlatformTimeZoneDetailParam `json:"params"`
}
