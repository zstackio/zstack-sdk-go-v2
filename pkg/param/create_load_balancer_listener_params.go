// Copyright (c) ZStack.io, Inc.

package param

// CreateLoadBalancerListenerDetailParam CreateLoadBalancerListener detail param
type CreateLoadBalancerListenerDetailParam struct {
	LoadBalancerUuid string `json:"loadBalancerUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	InstancePort int `json:"instancePort,omitempty"`
	LoadBalancerPort int `json:"loadBalancerPort" validate:"required"`
	Protocol string `json:"protocol,omitempty"`
	CertificateUuid string `json:"certificateUuid,omitempty"`
	HealthCheckProtocol string `json:"healthCheckProtocol,omitempty"`
	HealthCheckMethod string `json:"healthCheckMethod,omitempty"`
	HealthCheckURI string `json:"healthCheckURI,omitempty"`
	HealthCheckHttpCode string `json:"healthCheckHttpCode,omitempty"`
	AclStatus string `json:"aclStatus,omitempty"`
	AclUuids []string `json:"aclUuids,omitempty"`
	AclType string `json:"aclType,omitempty"`
	SecurityPolicyType string `json:"securityPolicyType,omitempty"`
	HttpVersions []string `json:"httpVersions,omitempty"`
	TcpProxyProtocol string `json:"tcpProxyProtocol,omitempty"`
	HttpCompressAlgos []string `json:"httpCompressAlgos,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateLoadBalancerListenerParam CreateLoadBalancerListener request param
type CreateLoadBalancerListenerParam struct {
	BaseParam
	Params CreateLoadBalancerListenerDetailParam `json:"params"`
}
