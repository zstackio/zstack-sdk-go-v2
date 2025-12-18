// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachAutoScalingTemplateFromGroup operates on AutoScalingTemplateFromGroup
func (cli *ZSClient) DetachAutoScalingTemplateFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/template/{templateUuid}/groups/{groupUuid}", uuid, string(deleteMode))
}
