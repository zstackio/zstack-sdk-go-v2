// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVipQos deletes VipQos
func (cli *ZSClient) DeleteVipQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/{uuid}/vip-qos", uuid, string(deleteMode))
}
