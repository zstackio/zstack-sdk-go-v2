// Copyright (c) ZStack.io, Inc.

package param

// CheckIpAvailabilityDetailParam CheckIpAvailability详细参数
type CheckIpAvailabilityDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"ip" validate:"required"` // 必填
	rest bool `json:"arpCheck,omitempty"`
	rest bool `json:"ipRangeCheck,omitempty"`
}

// CheckIpAvailabilityParam CheckIpAvailability请求参数
type CheckIpAvailabilityParam struct {
	BaseParam
	Params CheckIpAvailabilityDetailParam `json:"params"` // 详细参数
}

