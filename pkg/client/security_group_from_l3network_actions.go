// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachSecurityGroupFromL3Network 操作SecurityGroupFromL3Network
func (cli *ZSClient) DetachSecurityGroupFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}", uuid, string(deleteMode))
}

