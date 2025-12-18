// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteFlowMeter 删除FlowMeter
func (cli *ZSClient) DeleteFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/{uuid}", uuid, string(deleteMode))
}

