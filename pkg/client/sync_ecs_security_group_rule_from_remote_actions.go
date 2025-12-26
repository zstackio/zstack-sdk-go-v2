// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncEcsSecurityGroupRuleFromRemote operates on SyncEcsSecurityGroupRuleFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupRuleFromRemote(uuid string, params param.SyncEcsSecurityGroupRuleFromRemoteParam) (*view.SyncEcsSecurityGroupRuleFromRemoteEventView, error) {
	resp := view.SyncEcsSecurityGroupRuleFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group-rule/{uuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
