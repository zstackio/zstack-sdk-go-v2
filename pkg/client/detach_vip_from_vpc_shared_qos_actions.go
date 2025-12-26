// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachVipFromVpcSharedQos operates on VipFromVpcSharedQos
func (cli *ZSClient) DetachVipFromVpcSharedQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/sharedqos/{sharedQosUuid}/vips", uuid, string(deleteMode))
}
