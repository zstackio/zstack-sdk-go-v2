// Copyright (c) ZStack.io, Inc.

package param

// CreateVpnIpsecConfigDetailParam CreateVpnIpsecConfig detail param
type CreateVpnIpsecConfigDetailParam struct {
	Name string `json:"name" validate:"required"`
	Pfs string `json:"pfs,omitempty"`
	EncAlg string `json:"encAlg,omitempty"`
	AuthAlg string `json:"authAlg,omitempty"`
	Lifetime int `json:"lifetime,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpnIpsecConfigParam CreateVpnIpsecConfig request param
type CreateVpnIpsecConfigParam struct {
	BaseParam
	Params CreateVpnIpsecConfigDetailParam `json:"params"`
}
