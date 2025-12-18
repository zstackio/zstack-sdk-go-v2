// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncEcsSecurityGroupRuleFromRemote 操作SyncEcsSecurityGroupRuleFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupRuleFromRemote(uuid string, params param.SyncEcsSecurityGroupRuleFromRemoteParam) (*view.SyncEcsSecurityGroupRuleFromRemoteEventView, error) {
	resp := view.SyncEcsSecurityGroupRuleFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group-rule/{uuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

