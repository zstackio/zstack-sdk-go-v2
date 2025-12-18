// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncEcsSecurityGroupFromRemote operates on SyncEcsSecurityGroupFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupFromRemote(uuid string, params param.SyncEcsSecurityGroupFromRemoteParam) (*view.SyncEcsSecurityGroupFromRemoteEventView, error) {
	resp := view.SyncEcsSecurityGroupFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group/{ecsVpcUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
