// Copyright (c) ZStack.io, Inc.

package param

// AttachL3NetworksToIPsecConnectionDetailParam AttachL3NetworksToIPsecConnection详细参数
type AttachL3NetworksToIPsecConnectionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AttachL3NetworksToIPsecConnectionParam AttachL3NetworksToIPsecConnection请求参数
type AttachL3NetworksToIPsecConnectionParam struct {
	BaseParam
	Params AttachL3NetworksToIPsecConnectionDetailParam `json:"params"` // 详细参数
}

