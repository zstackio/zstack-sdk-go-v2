// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch 删除ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/connections/{uuid}", uuid, string(deleteMode))
}

