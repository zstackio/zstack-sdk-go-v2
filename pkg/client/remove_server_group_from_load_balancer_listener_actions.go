// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveServerGroupFromLoadBalancerListener removes ServerGroupFromLoadBalancerListener
func (cli *ZSClient) RemoveServerGroupFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/servergroups", uuid, string(deleteMode))
}
