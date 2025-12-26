// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSSOClient deletes SSOClient
func (cli *ZSClient) DeleteSSOClient(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/sso/client", uuid, string(deleteMode))
}
