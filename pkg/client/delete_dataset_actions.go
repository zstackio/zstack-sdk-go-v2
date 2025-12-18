// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteDataset deletes Dataset
func (cli *ZSClient) DeleteDataset(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/datasets/{uuid}", uuid, string(deleteMode))
}
