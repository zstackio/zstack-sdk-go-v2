// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSlbGroup deletes SlbGroup
func (cli *ZSClient) DeleteSlbGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/slb/group/{uuid}", uuid, string(deleteMode))
}
