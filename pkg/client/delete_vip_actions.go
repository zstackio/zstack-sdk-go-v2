// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVip deletes Vip
func (cli *ZSClient) DeleteVip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/{uuid}", uuid, string(deleteMode))
}
