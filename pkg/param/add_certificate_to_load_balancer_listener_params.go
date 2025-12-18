// Copyright (c) ZStack.io, Inc.

package param

// AddCertificateToLoadBalancerListenerDetailParam AddCertificateToLoadBalancerListener detail param
type AddCertificateToLoadBalancerListenerDetailParam struct {
	CertificateUuid string `json:"certificateUuid" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// AddCertificateToLoadBalancerListenerParam AddCertificateToLoadBalancerListener request param
type AddCertificateToLoadBalancerListenerParam struct {
	BaseParam
	Params AddCertificateToLoadBalancerListenerDetailParam `json:"params"`
}
