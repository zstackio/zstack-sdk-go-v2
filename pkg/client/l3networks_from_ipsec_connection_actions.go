// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachL3NetworksFromIPsecConnection 操作L3NetworksFromIPsecConnection
func (cli *ZSClient) DetachL3NetworksFromIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec/{uuid}/l3networks", uuid, string(deleteMode))
}

