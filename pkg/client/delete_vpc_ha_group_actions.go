// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcHaGroup deletes VpcHaGroup
func (cli *ZSClient) DeleteVpcHaGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpc/hagroups/{uuid}", uuid, string(deleteMode))
}
