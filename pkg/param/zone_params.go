// Copyright (c) ZStack.io, Inc.

package param

// GetZoneDetailParam GetZone详细参数
type GetZoneDetailParam struct {
	rest string `json:"uuid,omitempty"`
}

// GetZoneParam GetZone请求参数
type GetZoneParam struct {
	BaseParam
	Params GetZoneDetailParam `json:"params"` // 详细参数
}

