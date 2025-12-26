// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteFlowCollector deletes FlowCollector
func (cli *ZSClient) DeleteFlowCollector(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/collectors/{uuid}", uuid, string(deleteMode))
}
