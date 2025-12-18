// Copyright (c) ZStack.io, Inc.

package param

// AddRemoteCidrsToIPsecConnectionDetailParam AddRemoteCidrsToIPsecConnection详细参数
type AddRemoteCidrsToIPsecConnectionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"peerCidrs" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddRemoteCidrsToIPsecConnectionParam AddRemoteCidrsToIPsecConnection请求参数
type AddRemoteCidrsToIPsecConnectionParam struct {
	BaseParam
	Params AddRemoteCidrsToIPsecConnectionDetailParam `json:"params"` // 详细参数
}

