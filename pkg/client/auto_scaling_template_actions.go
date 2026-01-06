// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteAutoScalingTemplate deletes AutoScalingTemplate
func (cli *ZSClient) DeleteAutoScalingTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/template/{uuid}", uuid, string(deleteMode))
}
