// Copyright (c) ZStack.io, Inc.

package param

// CreateVpnIkeConfigDetailParam CreateVpnIkeConfig详细参数
type CreateVpnIkeConfigDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"psk" validate:"required"` // 必填
	rest string `json:"pfs,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"mode,omitempty"`
	rest string `json:"encAlg,omitempty"`
	rest string `json:"authAlg,omitempty"`
	rest int `json:"lifetime,omitempty"`
	rest string `json:"localIp" validate:"required"` // 必填
	rest string `json:"remoteIp" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVpnIkeConfigParam CreateVpnIkeConfig请求参数
type CreateVpnIkeConfigParam struct {
	BaseParam
	Params CreateVpnIkeConfigDetailParam `json:"params"` // 详细参数
}

