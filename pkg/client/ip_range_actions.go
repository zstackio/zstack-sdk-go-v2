// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteIpRange 删除IpRange
func (cli *ZSClient) DeleteIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/ip-ranges/{uuid}", uuid, string(deleteMode))
}

