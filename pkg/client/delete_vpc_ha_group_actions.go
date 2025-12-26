// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVpcHaGroup deletes VpcHaGroup
func (cli *ZSClient) DeleteVpcHaGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpc/hagroups/{uuid}", uuid, string(deleteMode))
}
