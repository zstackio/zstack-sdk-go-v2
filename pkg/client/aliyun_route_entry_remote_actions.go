// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunRouteEntryRemote 删除AliyunRouteEntryRemote
func (cli *ZSClient) DeleteAliyunRouteEntryRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/route-entry/{uuid}", uuid, string(deleteMode))
}

