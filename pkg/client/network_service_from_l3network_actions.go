// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachNetworkServiceFromL3Network 操作NetworkServiceFromL3Network
func (cli *ZSClient) DetachNetworkServiceFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/network-services", uuid, string(deleteMode))
}

