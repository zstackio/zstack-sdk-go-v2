// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveAccessControlListFromLoadBalancer removes AccessControlListFromLoadBalancer
func (cli *ZSClient) RemoveAccessControlListFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/access-control-lists", uuid, string(deleteMode))
}
