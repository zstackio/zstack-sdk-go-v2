// Copyright (c) ZStack.io, Inc.

package param

// DetachScsiLunFromHostDetailParam DetachScsiLunFromHost详细参数
type DetachScsiLunFromHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"hostUuid,omitempty"`
}

// DetachScsiLunFromHostParam DetachScsiLunFromHost请求参数
type DetachScsiLunFromHostParam struct {
	BaseParam
	Params DetachScsiLunFromHostDetailParam `json:"params"` // 详细参数
}

