// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVniRange 删除VniRange
func (cli *ZSClient) DeleteVniRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/vxlan-pool/vni-ranges/{uuid}", uuid, string(deleteMode))
}

