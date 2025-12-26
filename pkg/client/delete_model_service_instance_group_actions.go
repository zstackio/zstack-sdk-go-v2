// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModelServiceInstanceGroup deletes ModelServiceInstanceGroup
func (cli *ZSClient) DeleteModelServiceInstanceGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/instances/groups/{uuid}", uuid, string(deleteMode))
}
