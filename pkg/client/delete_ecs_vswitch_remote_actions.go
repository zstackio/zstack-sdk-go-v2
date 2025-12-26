// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEcsVSwitchRemote deletes EcsVSwitchRemote
func (cli *ZSClient) DeleteEcsVSwitchRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch/remote/{uuid}", uuid, string(deleteMode))
}
