// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteIdentityZoneInLocal deletes IdentityZoneInLocal
func (cli *ZSClient) DeleteIdentityZoneInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/identity-zone/{uuid}", uuid, string(deleteMode))
}
