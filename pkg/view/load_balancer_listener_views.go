// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerListenerInventoryView LoadBalancerListener
type LoadBalancerListenerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"loadBalancerUuid,omitempty"`
	rest int `json:"instancePort,omitempty"`
	rest int `json:"loadBalancerPort,omitempty"`
	rest string `json:"securityPolicyType,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest string `json:"serverGroupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []LoadBalancerListenerVmNicRefInventoryView `json:"vmNicRefs,omitempty"`
	rest []LoadBalancerListenerACLRefInventoryView `json:"aclRefs,omitempty"`
	rest []LoadBalancerListenerCertificateRefInventoryView `json:"certificateRefs,omitempty"`
	rest []LoadBalancerListenerServerGroupRefInventoryView `json:"serverGroupRefs,omitempty"`
}

