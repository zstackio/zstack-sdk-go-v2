// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteReservedIpRange 删除ReservedIpRange
func (cli *ZSClient) DeleteReservedIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/reserved-ip-ranges/{uuid}", uuid, string(deleteMode))
}

