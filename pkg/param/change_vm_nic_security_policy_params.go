// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicSecurityPolicyDetailParam ChangeVmNicSecurityPolicy detail param
type ChangeVmNicSecurityPolicyDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	IngressPolicy string `json:"ingressPolicy,omitempty"`
	EgressPolicy string `json:"egressPolicy,omitempty"`
}

// ChangeVmNicSecurityPolicyParam ChangeVmNicSecurityPolicy request param
type ChangeVmNicSecurityPolicyParam struct {
	BaseParam
	Params ChangeVmNicSecurityPolicyDetailParam `json:"params"`
}
