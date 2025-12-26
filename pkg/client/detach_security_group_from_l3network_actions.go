// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachSecurityGroupFromL3Network operates on SecurityGroupFromL3Network
func (cli *ZSClient) DetachSecurityGroupFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}", uuid, string(deleteMode))
}
