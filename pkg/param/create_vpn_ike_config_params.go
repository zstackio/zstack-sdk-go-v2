// Copyright (c) ZStack.io, Inc.

package param

// CreateVpnIkeConfigDetailParam CreateVpnIkeConfig detail param
type CreateVpnIkeConfigDetailParam struct {
	Name string `json:"name" validate:"required"`
	Psk string `json:"psk" validate:"required"`
	Pfs string `json:"pfs,omitempty"`
	Version string `json:"version,omitempty"`
	Mode string `json:"mode,omitempty"`
	EncAlg string `json:"encAlg,omitempty"`
	AuthAlg string `json:"authAlg,omitempty"`
	Lifetime int `json:"lifetime,omitempty"`
	LocalIp string `json:"localIp" validate:"required"`
	RemoteIp string `json:"remoteIp" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpnIkeConfigParam CreateVpnIkeConfig request param
type CreateVpnIkeConfigParam struct {
	BaseParam
	Params CreateVpnIkeConfigDetailParam `json:"params"`
}
