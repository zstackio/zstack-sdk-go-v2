// Copyright (c) ZStack.io, Inc.

package param

// CheckIpAvailabilityDetailParam CheckIpAvailability detail param
type CheckIpAvailabilityDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Ip string `json:"ip" validate:"required"`
	ArpCheck bool `json:"arpCheck,omitempty"`
	IpRangeCheck bool `json:"ipRangeCheck,omitempty"`
}

// CheckIpAvailabilityParam CheckIpAvailability request param
type CheckIpAvailabilityParam struct {
	BaseParam
	Params CheckIpAvailabilityDetailParam `json:"params"`
}
