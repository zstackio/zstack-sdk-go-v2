// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteResourceStackVmPortMonitor deletes ResourceStackVmPortMonitor
func (cli *ZSClient) DeleteResourceStackVmPortMonitor(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack/monitor/delvm", uuid, string(deleteMode))
}
