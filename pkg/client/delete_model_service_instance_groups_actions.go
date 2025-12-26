// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModelServiceInstanceGroups deletes ModelServiceInstanceGroups
func (cli *ZSClient) DeleteModelServiceInstanceGroups(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/instances/groups", uuid, string(deleteMode))
}
