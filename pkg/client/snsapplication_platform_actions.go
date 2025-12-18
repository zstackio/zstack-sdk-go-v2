// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteSNSApplicationPlatform 删除SNSApplicationPlatform
func (cli *ZSClient) DeleteSNSApplicationPlatform(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-platforms/{uuid}", uuid, string(deleteMode))
}

