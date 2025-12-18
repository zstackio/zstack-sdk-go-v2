// Copyright (c) ZStack.io, Inc.

package param

// RemoveCertificateFromLoadBalancerListenerDetailParam RemoveCertificateFromLoadBalancerListener detail param
type RemoveCertificateFromLoadBalancerListenerDetailParam struct {
	CertificateUuid string `json:"certificateUuid" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// RemoveCertificateFromLoadBalancerListenerParam RemoveCertificateFromLoadBalancerListener request param
type RemoveCertificateFromLoadBalancerListenerParam struct {
	BaseParam
	Params RemoveCertificateFromLoadBalancerListenerDetailParam `json:"params"`
}
