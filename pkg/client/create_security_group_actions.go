// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSecurityGroup creates SecurityGroup
func (cli *ZSClient) CreateSecurityGroup(params param.CreateSecurityGroupParam) (*view.CreateSecurityGroupEventView, error) {
	resp := view.CreateSecurityGroupEventView{}
	if err := cli.Post("v1/security-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
