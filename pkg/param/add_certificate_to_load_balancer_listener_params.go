// Copyright (c) ZStack.io, Inc.

package param

// AddCertificateToLoadBalancerListenerDetailParam AddCertificateToLoadBalancerListener详细参数
type AddCertificateToLoadBalancerListenerDetailParam struct {
	rest string `json:"certificateUuid" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
}

// AddCertificateToLoadBalancerListenerParam AddCertificateToLoadBalancerListener请求参数
type AddCertificateToLoadBalancerListenerParam struct {
	BaseParam
	Params AddCertificateToLoadBalancerListenerDetailParam `json:"params"` // 详细参数
}

