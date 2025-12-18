// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteUserProxyConfig deletes UserProxyConfig
func (cli *ZSClient) DeleteUserProxyConfig(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/user-proxy-configs/{uuid}", uuid, string(deleteMode))
}
