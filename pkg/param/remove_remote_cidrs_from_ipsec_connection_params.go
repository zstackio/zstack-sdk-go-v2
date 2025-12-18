// Copyright (c) ZStack.io, Inc.

package param

// RemoveRemoteCidrsFromIPsecConnectionDetailParam RemoveRemoteCidrsFromIPsecConnection详细参数
type RemoveRemoteCidrsFromIPsecConnectionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"peerCidrs" validate:"required"` // 必填
}

// RemoveRemoteCidrsFromIPsecConnectionParam RemoveRemoteCidrsFromIPsecConnection请求参数
type RemoveRemoteCidrsFromIPsecConnectionParam struct {
	BaseParam
	Params RemoveRemoteCidrsFromIPsecConnectionDetailParam `json:"params"` // 详细参数
}

