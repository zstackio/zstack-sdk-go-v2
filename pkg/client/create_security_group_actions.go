// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSecurityGroup creates SecurityGroup
func (cli *ZSClient) CreateSecurityGroup(params param.CreateSecurityGroupParam) (*view.CreateSecurityGroupEventView, error) {
	resp := view.CreateSecurityGroupEventView{}
	if err := cli.Post("v1/security-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
