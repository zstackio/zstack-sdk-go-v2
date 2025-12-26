// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLoadBalancerServerGroup deletes LoadBalancerServerGroup
func (cli *ZSClient) DeleteLoadBalancerServerGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/servergroups/{uuid}", uuid, string(deleteMode))
}
