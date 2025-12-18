// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteZone deletes Zone
func (cli *ZSClient) DeleteZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{uuid}", uuid, string(deleteMode))
}
