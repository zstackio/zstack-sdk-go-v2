// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePolicy deletes Policy
func (cli *ZSClient) DeletePolicy(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/policies/{uuid}", uuid, string(deleteMode))
}
