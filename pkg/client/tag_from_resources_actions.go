// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachTagFromResources 操作TagFromResources
func (cli *ZSClient) DetachTagFromResources(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags/{tagUuid}/resources", uuid, string(deleteMode))
}

