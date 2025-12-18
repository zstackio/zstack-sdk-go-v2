// Copyright (c) ZStack.io, Inc.

package param

// ChangeLoadBalancerBackendServerDetailParam ChangeLoadBalancerBackendServer详细参数
type ChangeLoadBalancerBackendServerDetailParam struct {
	rest string `json:"serverGroupUuid" validate:"required"` // 必填
	rest []interface{} `json:"vmNics,omitempty"`
	rest []interface{} `json:"servers,omitempty"`
}

// ChangeLoadBalancerBackendServerParam ChangeLoadBalancerBackendServer请求参数
type ChangeLoadBalancerBackendServerParam struct {
	BaseParam
	Params ChangeLoadBalancerBackendServerDetailParam `json:"params"` // 详细参数
}

