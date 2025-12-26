// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveVRouterNetworksFromFlowMeter removes VRouterNetworksFromFlowMeter
func (cli *ZSClient) RemoveVRouterNetworksFromFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/networks", uuid, string(deleteMode))
}
