// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteEcsVSwitchInLocal deletes EcsVSwitchInLocal
func (cli *ZSClient) DeleteEcsVSwitchInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch/{uuid}", uuid, string(deleteMode))
}
