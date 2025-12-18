// Copyright (c) ZStack.io, Inc.

package param

// AddProxyToResourceDetailParam AddProxyToResource detail param
type AddProxyToResourceDetailParam struct {
	ProxyUuid string `json:"proxyUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// AddProxyToResourceParam AddProxyToResource request param
type AddProxyToResourceParam struct {
	BaseParam
	Params AddProxyToResourceDetailParam `json:"params"`
}
