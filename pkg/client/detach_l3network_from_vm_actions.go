// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachL3NetworkFromVm operates on L3NetworkFromVm
func (cli *ZSClient) DetachL3NetworkFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/nics/{vmNicUuid}", uuid, string(deleteMode))
}
