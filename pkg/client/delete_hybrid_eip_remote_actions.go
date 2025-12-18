// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteHybridEipRemote deletes HybridEipRemote
func (cli *ZSClient) DeleteHybridEipRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/eip/{uuid}/remote", uuid, string(deleteMode))
}
