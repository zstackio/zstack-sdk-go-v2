// Copyright (c) ZStack.io, Inc.

package param

// AckEventDataDetailParam AckEventData detail param
type AckEventDataDetailParam struct {
	EventSubscriptionUuid string `json:"eventSubscriptionUuid" validate:"required"`
	AlertDataUuid string `json:"alertDataUuid" validate:"required"`
	DataType string `json:"dataType" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	AckPeriodSec int `json:"ackPeriodSec" validate:"required"`
}

// AckEventDataParam AckEventData request param
type AckEventDataParam struct {
	BaseParam
	Params AckEventDataDetailParam `json:"params"`
}
