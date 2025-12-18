// Copyright (c) ZStack.io, Inc.

package param

// CreateVipDetailParam CreateVip详细参数
type CreateVipDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"allocatorStrategy,omitempty"`
	rest string `json:"ipRangeUuid,omitempty"`
	rest string `json:"requiredIp,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVipParam CreateVip请求参数
type CreateVipParam struct {
	BaseParam
	Params CreateVipDetailParam `json:"params"` // 详细参数
}

