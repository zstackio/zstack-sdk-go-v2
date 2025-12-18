// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteSSOClient deletes SSOClient
func (cli *ZSClient) DeleteSSOClient(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/sso/client", uuid, string(deleteMode))
}
