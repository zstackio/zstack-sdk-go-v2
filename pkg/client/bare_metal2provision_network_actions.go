// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteBareMetal2ProvisionNetwork 删除BareMetal2ProvisionNetwork
func (cli *ZSClient) DeleteBareMetal2ProvisionNetwork(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/provision-networks/{uuid}", uuid, string(deleteMode))
}

