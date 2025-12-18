// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVxlanL2Network 删除VxlanL2Network
func (cli *ZSClient) DeleteVxlanL2Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/vxlan/{uuid}", uuid, string(deleteMode))
}

