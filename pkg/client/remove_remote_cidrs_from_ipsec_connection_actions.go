// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveRemoteCidrsFromIPsecConnection 操作RemoveRemoteCidrsFromIPsecConnection
func (cli *ZSClient) RemoveRemoteCidrsFromIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec/{uuid}/remote-cidrs", uuid, string(deleteMode))
}

