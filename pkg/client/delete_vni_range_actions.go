// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVniRange deletes VniRange
func (cli *ZSClient) DeleteVniRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/vxlan-pool/vni-ranges/{uuid}", uuid, string(deleteMode))
}
