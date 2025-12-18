// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicSecurityPolicyDetailParam ChangeVmNicSecurityPolicy详细参数
type ChangeVmNicSecurityPolicyDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"ingressPolicy,omitempty"`
	rest string `json:"egressPolicy,omitempty"`
}

// ChangeVmNicSecurityPolicyParam ChangeVmNicSecurityPolicy请求参数
type ChangeVmNicSecurityPolicyParam struct {
	BaseParam
	Params ChangeVmNicSecurityPolicyDetailParam `json:"params"` // 详细参数
}

