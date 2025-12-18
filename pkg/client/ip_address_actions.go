// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteIpAddress 删除IpAddress
func (cli *ZSClient) DeleteIpAddress(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/ip-address", uuid, string(deleteMode))
}

