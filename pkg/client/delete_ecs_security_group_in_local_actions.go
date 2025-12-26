// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsSecurityGroupInLocal deletes EcsSecurityGroupInLocal
func (cli *ZSClient) DeleteEcsSecurityGroupInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group/{uuid}", uuid, string(deleteMode))
}
