// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SecurityGroupRuleInventoryView SecurityGroupRule
type SecurityGroupRuleInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"securityGroupUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest string `json:"state,omitempty"`
	rest int `json:"priority,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"srcIpRange,omitempty"`
	rest string `json:"dstIpRange,omitempty"`
	rest string `json:"srcPortRange,omitempty"`
	rest string `json:"dstPortRange,omitempty"`
	rest string `json:"action,omitempty"`
	rest string `json:"remoteSecurityGroupUuid,omitempty"`
	rest string `json:"allowedCidr,omitempty"`
	rest int `json:"startPort,omitempty"`
	rest int `json:"endPort,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

