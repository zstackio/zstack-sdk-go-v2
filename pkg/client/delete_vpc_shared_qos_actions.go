// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVpcSharedQos deletes VpcSharedQos
func (cli *ZSClient) DeleteVpcSharedQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/sharedqos/{uuid}", uuid, string(deleteMode))
}
