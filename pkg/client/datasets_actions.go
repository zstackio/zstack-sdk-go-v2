// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteDatasets 删除Datasets
func (cli *ZSClient) DeleteDatasets(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/datasets", uuid, string(deleteMode))
}

