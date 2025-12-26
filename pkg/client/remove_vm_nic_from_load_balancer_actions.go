// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveVmNicFromLoadBalancer removes VmNicFromLoadBalancer
func (cli *ZSClient) RemoveVmNicFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics", uuid, string(deleteMode))
}
