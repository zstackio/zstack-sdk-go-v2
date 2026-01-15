// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateLoadBalancerParamDetail CreateLoadBalancer detail param
type CreateLoadBalancerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Ipv6VipUuid string `json:"ipv6VipUuid,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateLoadBalancerParam CreateLoadBalancer request param
type CreateLoadBalancerParam struct {
	BaseParam
	Params CreateLoadBalancerParamDetail `json:"createLoadBalancer"`
}
// UpdateLoadBalancerParamDetail UpdateLoadBalancer detail param
type UpdateLoadBalancerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateLoadBalancerParam UpdateLoadBalancer request param
type UpdateLoadBalancerParam struct {
	BaseParam
	Params UpdateLoadBalancerParamDetail `json:"updateLoadBalancer"`
}
// DeleteLoadBalancerParamDetail DeleteLoadBalancer detail param
type DeleteLoadBalancerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteLoadBalancerParam DeleteLoadBalancer request param
type DeleteLoadBalancerParam struct {
	BaseParam
	Params DeleteLoadBalancerParamDetail `json:"deleteLoadBalancer"`
}
// RefreshLoadBalancerParamDetail RefreshLoadBalancer detail param
type RefreshLoadBalancerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RefreshLoadBalancerParam RefreshLoadBalancer request param
type RefreshLoadBalancerParam struct {
	BaseParam
	Params RefreshLoadBalancerParamDetail `json:"refreshLoadBalancer"`
}
