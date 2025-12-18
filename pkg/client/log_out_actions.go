// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LogOut operates on LogOut
func (cli *ZSClient) LogOut(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/sessions/{sessionUuid}", uuid, string(deleteMode))
}
