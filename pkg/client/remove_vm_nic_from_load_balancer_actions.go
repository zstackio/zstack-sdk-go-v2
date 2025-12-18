// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveVmNicFromLoadBalancer 操作RemoveVmNicFromLoadBalancer
func (cli *ZSClient) RemoveVmNicFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics", uuid, string(deleteMode))
}

