// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteFlowMeter deletes FlowMeter
func (cli *ZSClient) DeleteFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/{uuid}", uuid, string(deleteMode))
}
