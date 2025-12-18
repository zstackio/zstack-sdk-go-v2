// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachEip operates on Eip
func (cli *ZSClient) DetachEip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/eips/{uuid}/vm-instances/nics", uuid, string(deleteMode))
}
