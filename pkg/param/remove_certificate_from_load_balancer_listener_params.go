// Copyright (c) ZStack.io, Inc.

package param

// RemoveCertificateFromLoadBalancerListenerDetailParam RemoveCertificateFromLoadBalancerListener详细参数
type RemoveCertificateFromLoadBalancerListenerDetailParam struct {
	rest string `json:"certificateUuid" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
}

// RemoveCertificateFromLoadBalancerListenerParam RemoveCertificateFromLoadBalancerListener请求参数
type RemoveCertificateFromLoadBalancerListenerParam struct {
	BaseParam
	Params RemoveCertificateFromLoadBalancerListenerDetailParam `json:"params"` // 详细参数
}

