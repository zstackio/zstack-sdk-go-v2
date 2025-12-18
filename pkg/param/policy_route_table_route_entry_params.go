// Copyright (c) ZStack.io, Inc.

package param

// CreatePolicyRouteTableRouteEntryDetailParam CreatePolicyRouteTableRouteEntry详细参数
type CreatePolicyRouteTableRouteEntryDetailParam struct {
	rest string `json:"tableUuid" validate:"required"` // 必填
	rest string `json:"destinationCidr" validate:"required"` // 必填
	rest string `json:"nextHopIp" validate:"required"` // 必填
	rest int `json:"distance,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteTableRouteEntryParam CreatePolicyRouteTableRouteEntry请求参数
type CreatePolicyRouteTableRouteEntryParam struct {
	BaseParam
	Params CreatePolicyRouteTableRouteEntryDetailParam `json:"params"` // 详细参数
}

