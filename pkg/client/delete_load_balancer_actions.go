// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLoadBalancer deletes LoadBalancer
func (cli *ZSClient) DeleteLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/{uuid}", uuid, string(deleteMode))
}
