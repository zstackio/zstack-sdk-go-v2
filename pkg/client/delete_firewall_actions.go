// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteFirewall deletes Firewall
func (cli *ZSClient) DeleteFirewall(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/{uuid}", uuid, string(deleteMode))
}
