// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteModelServiceInstanceGroups 删除ModelServiceInstanceGroups
func (cli *ZSClient) DeleteModelServiceInstanceGroups(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/instances/groups", uuid, string(deleteMode))
}

