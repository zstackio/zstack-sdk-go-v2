// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAllEcsInstancesFromDataCenter deletes AllEcsInstancesFromDataCenter
func (cli *ZSClient) DeleteAllEcsInstancesFromDataCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/dc-ecs/{uuid}", uuid, string(deleteMode))
}
