// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsSecurityGroupInventoryView EcsSecurityGroup
type EcsSecurityGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	EcsVpcUuid string `json:"ecsVpcUuid,omitempty"`
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

