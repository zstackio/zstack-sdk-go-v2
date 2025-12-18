// Copyright (c) ZStack.io, Inc.

package param

// CheckVipPortAvailabilityDetailParam CheckVipPortAvailability详细参数
type CheckVipPortAvailabilityDetailParam struct {
	rest string `json:"vipUuid" validate:"required"` // 必填
	rest int `json:"port" validate:"required"` // 必填
	rest string `json:"protocolType" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// CheckVipPortAvailabilityParam CheckVipPortAvailability请求参数
type CheckVipPortAvailabilityParam struct {
	BaseParam
	Params CheckVipPortAvailabilityDetailParam `json:"params"` // 详细参数
}

