// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePolicyRouteTable deletes PolicyRouteTable
func (cli *ZSClient) DeletePolicyRouteTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/tables/{uuid}", uuid, string(deleteMode))
}
