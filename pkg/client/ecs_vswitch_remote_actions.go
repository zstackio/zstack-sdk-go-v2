// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteEcsVSwitchRemote 删除EcsVSwitchRemote
func (cli *ZSClient) DeleteEcsVSwitchRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch/remote/{uuid}", uuid, string(deleteMode))
}

