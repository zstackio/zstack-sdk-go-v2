// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEcsSecurityGroupRuleRemote creates EcsSecurityGroupRuleRemote
func (cli *ZSClient) CreateEcsSecurityGroupRuleRemote(params param.CreateEcsSecurityGroupRuleRemoteParam) (*view.CreateEcsSecurityGroupRuleRemoteEventView, error) {
	resp := view.CreateEcsSecurityGroupRuleRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/security-group-rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
