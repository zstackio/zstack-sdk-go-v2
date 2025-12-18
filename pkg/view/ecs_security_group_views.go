// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EcsSecurityGroupInventoryView EcsSecurityGroup
type EcsSecurityGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"ecsVpcUuid,omitempty"`
	rest string `json:"securityGroupId,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

