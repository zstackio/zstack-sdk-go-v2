// Copyright (c) ZStack.io, Inc.

package param

// ChangeLoadBalancerListenerDetailParam ChangeLoadBalancerListener detail param
type ChangeLoadBalancerListenerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ConnectionIdleTimeout int `json:"connectionIdleTimeout,omitempty"`
	MaxConnection int `json:"maxConnection,omitempty"`
	BalancerAlgorithm string `json:"balancerAlgorithm,omitempty"`
	HealthCheckTarget string `json:"healthCheckTarget,omitempty"`
	HealthyThreshold int `json:"healthyThreshold,omitempty"`
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty"`
	HealthCheckInterval int `json:"healthCheckInterval,omitempty"`
	HealthCheckProtocol string `json:"healthCheckProtocol,omitempty"`
	HealthCheckMethod string `json:"healthCheckMethod,omitempty"`
	HealthCheckURI string `json:"healthCheckURI,omitempty"`
	HealthCheckHttpCode string `json:"healthCheckHttpCode,omitempty"`
	AclStatus string `json:"aclStatus,omitempty"`
	SecurityPolicyType string `json:"securityPolicyType,omitempty"`
	Nbprocess int `json:"nbprocess,omitempty"`
	HttpMode string `json:"httpMode,omitempty"`
	SessionPersistence string `json:"sessionPersistence,omitempty"`
	SessionIdleTimeout int `json:"sessionIdleTimeout,omitempty"`
	CookieName string `json:"cookieName,omitempty"`
	HttpVersions []string `json:"httpVersions,omitempty"`
	HttpRedirectHttps string `json:"httpRedirectHttps,omitempty"`
	RedirectPort int `json:"redirectPort,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
	TcpProxyProtocol string `json:"tcpProxyProtocol,omitempty"`
	HttpCompressAlgos []string `json:"httpCompressAlgos,omitempty"`
}

// ChangeLoadBalancerListenerParam ChangeLoadBalancerListener request param
type ChangeLoadBalancerListenerParam struct {
	BaseParam
	Params ChangeLoadBalancerListenerDetailParam `json:"params"`
}
