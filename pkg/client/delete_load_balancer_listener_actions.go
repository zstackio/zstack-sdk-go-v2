// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLoadBalancerListener deletes LoadBalancerListener
func (cli *ZSClient) DeleteLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{uuid}", uuid, string(deleteMode))
}
