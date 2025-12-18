// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveAccessControlListFromLoadBalancer 操作RemoveAccessControlListFromLoadBalancer
func (cli *ZSClient) RemoveAccessControlListFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/access-control-lists", uuid, string(deleteMode))
}

