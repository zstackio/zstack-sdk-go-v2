// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExpungeImage operates on Image
func (cli *ZSClient) ExpungeImage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/images/{imageUuid}/actions", uuid, string(deleteMode))
}
