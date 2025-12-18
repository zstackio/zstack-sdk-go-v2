// Copyright (c) ZStack.io, Inc.

package param

// CreateVpnIpsecConfigDetailParam CreateVpnIpsecConfig详细参数
type CreateVpnIpsecConfigDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"pfs,omitempty"`
	rest string `json:"encAlg,omitempty"`
	rest string `json:"authAlg,omitempty"`
	rest int `json:"lifetime,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVpnIpsecConfigParam CreateVpnIpsecConfig请求参数
type CreateVpnIpsecConfigParam struct {
	BaseParam
	Params CreateVpnIpsecConfigDetailParam `json:"params"` // 详细参数
}

