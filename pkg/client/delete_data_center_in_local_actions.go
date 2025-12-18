// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteDataCenterInLocal deletes DataCenterInLocal
func (cli *ZSClient) DeleteDataCenterInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/data-center/{uuid}", uuid, string(deleteMode))
}
