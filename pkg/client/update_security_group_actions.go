// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSecurityGroup updates SecurityGroup
func (cli *ZSClient) UpdateSecurityGroup(uuid string, params param.UpdateSecurityGroupParam) (*view.UpdateSecurityGroupEventView, error) {
	resp := view.UpdateSecurityGroupEventView{}
	if err := cli.Put("v1/security-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
