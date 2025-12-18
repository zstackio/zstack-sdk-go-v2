// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteBaremetalChassis deletes BaremetalChassis
func (cli *ZSClient) DeleteBaremetalChassis(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/chassis/{uuid}", uuid, string(deleteMode))
}
