// Copyright (c) ZStack.io, Inc.

package param

// AddProxyToResourceDetailParam AddProxyToResource详细参数
type AddProxyToResourceDetailParam struct {
	rest string `json:"proxyUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid" validate:"required"` // 必填
}

// AddProxyToResourceParam AddProxyToResource请求参数
type AddProxyToResourceParam struct {
	BaseParam
	Params AddProxyToResourceDetailParam `json:"params"` // 详细参数
}

