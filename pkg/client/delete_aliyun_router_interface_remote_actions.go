// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunRouterInterfaceRemote deletes AliyunRouterInterfaceRemote
func (cli *ZSClient) DeleteAliyunRouterInterfaceRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface/remote/{uuid}", uuid, string(deleteMode))
}
