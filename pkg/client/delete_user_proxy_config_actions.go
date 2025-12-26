// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteUserProxyConfig deletes UserProxyConfig
func (cli *ZSClient) DeleteUserProxyConfig(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/user-proxy-configs/{uuid}", uuid, string(deleteMode))
}
