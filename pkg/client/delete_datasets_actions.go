// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteDatasets deletes Datasets
func (cli *ZSClient) DeleteDatasets(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/datasets", uuid, string(deleteMode))
}
