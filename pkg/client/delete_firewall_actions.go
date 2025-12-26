// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteFirewall deletes Firewall
func (cli *ZSClient) DeleteFirewall(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/{uuid}", uuid, string(deleteMode))
}
