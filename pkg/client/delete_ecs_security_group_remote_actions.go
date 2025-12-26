// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsSecurityGroupRemote deletes EcsSecurityGroupRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group/remote/{uuid}", uuid, string(deleteMode))
}
