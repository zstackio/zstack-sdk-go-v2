// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVxlanL2Network deletes VxlanL2Network
func (cli *ZSClient) DeleteVxlanL2Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/vxlan/{uuid}", uuid, string(deleteMode))
}
