// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAllEcsInstancesFromDataCenter 删除AllEcsInstancesFromDataCenter
func (cli *ZSClient) DeleteAllEcsInstancesFromDataCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/dc-ecs/{uuid}", uuid, string(deleteMode))
}

