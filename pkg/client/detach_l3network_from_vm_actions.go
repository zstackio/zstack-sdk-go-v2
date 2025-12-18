// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachL3NetworkFromVm operates on L3NetworkFromVm
func (cli *ZSClient) DetachL3NetworkFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/nics/{vmNicUuid}", uuid, string(deleteMode))
}
