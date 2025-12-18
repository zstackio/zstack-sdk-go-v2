// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteBonding deletes Bonding
func (cli *ZSClient) DeleteBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/bondings/{uuid}", uuid, string(deleteMode))
}
