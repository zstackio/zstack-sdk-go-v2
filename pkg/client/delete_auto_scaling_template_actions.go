// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAutoScalingTemplate deletes AutoScalingTemplate
func (cli *ZSClient) DeleteAutoScalingTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/template/{uuid}", uuid, string(deleteMode))
}
