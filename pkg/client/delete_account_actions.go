// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAccount deletes Account
func (cli *ZSClient) DeleteAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/{uuid}", uuid, string(deleteMode))
}
