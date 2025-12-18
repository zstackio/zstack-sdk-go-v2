// Copyright (c) ZStack.io, Inc.

package param

// AddIpRangeByNetworkCidrDetailParam AddIpRangeByNetworkCidr详细参数
type AddIpRangeByNetworkCidrDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"networkCidr" validate:"required"` // 必填
	rest string `json:"gateway,omitempty"`
	rest string `json:"ipRangeType,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddIpRangeByNetworkCidrParam AddIpRangeByNetworkCidr请求参数
type AddIpRangeByNetworkCidrParam struct {
	BaseParam
	Params AddIpRangeByNetworkCidrDetailParam `json:"params"` // 详细参数
}

