// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePublishApp deletes PublishApp
func (cli *ZSClient) DeletePublishApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/app/{uuid}", uuid, string(deleteMode))
}
