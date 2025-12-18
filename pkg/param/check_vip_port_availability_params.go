// Copyright (c) ZStack.io, Inc.

package param

// CheckVipPortAvailabilityDetailParam CheckVipPortAvailability detail param
type CheckVipPortAvailabilityDetailParam struct {
	VipUuid string `json:"vipUuid" validate:"required"`
	Port int `json:"port" validate:"required"`
	ProtocolType string `json:"protocolType" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// CheckVipPortAvailabilityParam CheckVipPortAvailability request param
type CheckVipPortAvailabilityParam struct {
	BaseParam
	Params CheckVipPortAvailabilityDetailParam `json:"params"`
}
