// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateEcsSecurityGroupRemote creates EcsSecurityGroupRemote
func (cli *ZSClient) CreateEcsSecurityGroupRemote(params param.CreateEcsSecurityGroupRemoteParam) (*view.CreateEcsSecurityGroupRemoteEventView, error) {
	resp := view.CreateEcsSecurityGroupRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/security-group/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
