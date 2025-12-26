// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachAutoScalingTemplateFromGroup operates on AutoScalingTemplateFromGroup
func (cli *ZSClient) DetachAutoScalingTemplateFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/template/{templateUuid}/groups/{groupUuid}", uuid, string(deleteMode))
}
