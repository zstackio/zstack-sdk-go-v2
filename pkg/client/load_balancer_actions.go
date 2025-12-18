// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteLoadBalancer 删除LoadBalancer
func (cli *ZSClient) DeleteLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/{uuid}", uuid, string(deleteMode))
}

