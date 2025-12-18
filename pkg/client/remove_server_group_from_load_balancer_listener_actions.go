// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveServerGroupFromLoadBalancerListener 操作RemoveServerGroupFromLoadBalancerListener
func (cli *ZSClient) RemoveServerGroupFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/servergroups", uuid, string(deleteMode))
}

