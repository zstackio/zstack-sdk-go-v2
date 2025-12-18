// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteFlowCollector 删除FlowCollector
func (cli *ZSClient) DeleteFlowCollector(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/collectors/{uuid}", uuid, string(deleteMode))
}

