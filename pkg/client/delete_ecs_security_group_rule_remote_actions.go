// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsSecurityGroupRuleRemote deletes EcsSecurityGroupRuleRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRuleRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group-rule/remote/{uuid}", uuid, string(deleteMode))
}
