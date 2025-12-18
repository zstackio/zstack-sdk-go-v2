// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteEcsVSwitchInLocal 删除EcsVSwitchInLocal
func (cli *ZSClient) DeleteEcsVSwitchInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch/{uuid}", uuid, string(deleteMode))
}

