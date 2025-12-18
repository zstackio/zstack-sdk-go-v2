// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerListenerInventoryView LoadBalancerListener
type LoadBalancerListenerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid,omitempty"`
	InstancePort int `json:"instancePort,omitempty"`
	LoadBalancerPort int `json:"loadBalancerPort,omitempty"`
	SecurityPolicyType string `json:"securityPolicyType,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	VmNicRefs []LoadBalancerListenerVmNicRefInventoryView `json:"vmNicRefs,omitempty"`
	AclRefs []LoadBalancerListenerACLRefInventoryView `json:"aclRefs,omitempty"`
	CertificateRefs []LoadBalancerListenerCertificateRefInventoryView `json:"certificateRefs,omitempty"`
	ServerGroupRefs []LoadBalancerListenerServerGroupRefInventoryView `json:"serverGroupRefs,omitempty"`
}

