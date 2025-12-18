// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAccessControlRule deletes AccessControlRule
func (cli *ZSClient) DeleteAccessControlRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/login-control/access-control/rules/{uuid}", uuid, string(deleteMode))
}
