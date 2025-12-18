// Copyright (c) ZStack.io, Inc.

package param

// AckEventDataDetailParam AckEventData详细参数
type AckEventDataDetailParam struct {
	rest string `json:"eventSubscriptionUuid" validate:"required"` // 必填
	rest string `json:"alertDataUuid" validate:"required"` // 必填
	rest string `json:"dataType" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest int `json:"ackPeriodSec" validate:"required"` // 必填
}

// AckEventDataParam AckEventData请求参数
type AckEventDataParam struct {
	BaseParam
	Params AckEventDataDetailParam `json:"params"` // 详细参数
}

