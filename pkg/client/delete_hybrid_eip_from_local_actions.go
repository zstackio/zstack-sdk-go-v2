// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteHybridEipFromLocal deletes HybridEipFromLocal
func (cli *ZSClient) DeleteHybridEipFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/eip/{uuid}", uuid, string(deleteMode))
}
